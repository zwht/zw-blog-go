package http

// 接口权限处理
import (
	. "../../datamodels"
	. "../../models"
	. "../../tools"
	"encoding/json"
	"github.com/kataras/iris"
	"strings"
	"sync"
)

type Context struct {
	iris.Context
	User User
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original
	tokenString := ctx.Request().Header.Get("Authorization")
	//如果url有token，优先使用url
	urlToken := ctx.Params().Get("token")
	if urlToken != "" {
		tokenString = urlToken
	}
	if tokenString != "" {
		_, user := GetJwt(tokenString)
		var empList User
		_ = json.Unmarshal(user, &empList)
		ctx.User = empList
	}
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

func Permission(h func(*Context), roles string) iris.Handler {
	return func(original iris.Context) {
		//处理ctx，把操作user数据绑定在ctx
		ctx := acquire(original)
		passe := false

		urlRoles := strings.Split(roles, ",")
		if roles == "" {
			passe = true
		} else {
			userRoles := strings.Split(ctx.User.Roles, ",")
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
