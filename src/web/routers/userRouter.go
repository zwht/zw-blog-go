package routers

import (
	. "../../config"
	. "../controllers"
	"github.com/kataras/iris"
)

//user model
func UserRouter(app *iris.Application) {
	routes := app.Party("/v1/user", func(ctx iris.Context) {

		ctx.Next()
	})
	{
		routes.Post("/add", Permission(UserCreate, "1001"))
		routes.Post("/update", UserUpdateCtr)
		routes.Get("/del/{id:string}", Permission(UserDeleteById, "1001"))
		routes.Get("/getById/{id:string}", UserGetById)
		routes.Post("/list/{pageNum:int}/{pageSize:int}", Permission(UserGetList, "1001"))
		routes.Post("/login", Login)
	}

}
