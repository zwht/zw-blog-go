package config

import (
	. "../datamodels"
	. "../models"
	"encoding/json"
	"github.com/kataras/iris"
	"strings"
	"sync"
)

// Context is our custom context.
// Let's implement a context which will give us access
// to the client's Session with a trivial `ctx.Session()` call.
type Context struct {
	iris.Context
	user User
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original
	tokenString := ctx.Request().Header.Get("Authorization")
	if tokenString != "" {
		_, user := GetJwt(tokenString)
		var empList User
		_ = json.Unmarshal(user, &empList)
		ctx.user = empList
	}
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

func Permission(h func(iris.Context), roles string) iris.Handler {
	return func(original iris.Context) {
		//处理ctx，把操作user数据绑定在ctx
		ctx := acquire(original)
		passe := false

		urlRoles := strings.Split(roles, ",")
		if len(urlRoles) == 0 {
			passe = true
		} else {
			userRoles := strings.Split(ctx.user.Roles, ",")
			for _, userR := range userRoles {
				for _, urlR := range urlRoles {
					if userR == urlR {
						passe = true
					}
				}
			}
		}

		if passe {
			h(ctx)
			release(ctx)
		} else {
			result := Result{}
			result.Code = 402
			result.Msg = "接口没有权限"
			ctx.JSON(result)
		}
	}
}
