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
		routes.Post("/add", UserCreate)
		routes.Post("/update", UserUpdate)
		routes.Get("/del/{id:string}", UserDeleteById)
		routes.Get("/getById/{id:string}", UserGetById)
		routes.Post("/list/{pageSize:int}/{pageNum:int}", UserGetList)
		routes.Post("/login", Login)
	}

}
