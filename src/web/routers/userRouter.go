package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//user model
func UserRouter(app *iris.Application) {
	routes := app.Party("/v1/user", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", CreateUser)
		routes.Post("/update", UpdateUser)
		routes.Get("/del/{id:int}", DeleteUserById)
		routes.Get("/getById/{id:int}", GetUserById)
		routes.Post("/list", GetUserList)
		routes.Post("/login", Login)
	}

}
