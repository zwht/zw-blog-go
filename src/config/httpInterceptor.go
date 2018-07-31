package config

import (
	. "../datamodels"
	"github.com/kataras/iris"
	"strings"
)

var noTokenMap = [...]string{
	"/v1/user/login",
	"/v1/code/list",
}

func HttpInterceptor(ctx iris.Context) {
	var isNext = true
	requestPath := ctx.Path()
	if strings.Contains(requestPath, "/v1") {
		var i int
		isNext = false
		for i = 0; i < len(noTokenMap); i++ {
			if strings.Contains(requestPath, noTokenMap[i]) {
				isNext = true
			}
		}
		if !isNext {
			tokenString := ctx.Request().Header.Get("Authorization")
			if tokenString != "" {
				isNext, _ = GetJwt(tokenString)
			}
			// var empList User
			// err3 := json.Unmarshal(tokenStr33, &empList)
			// if err3 != nil {
			// 	fmt.Printf(err3.Error())
			// }
		}
	}
	if isNext {
		ctx.Next()
	} else {
		result := Result{}
		result.Code = 401
		result.Msg = "没有登录或者登录过期"
		ctx.JSON(result)
	}

}
