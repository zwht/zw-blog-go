package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func UserGroupRouter(app *iris.Application) {
	routes := app.Party("/v1/user/group", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", UserGroupCreate)
		routes.Post("/update", UserGroupUpdate)
		routes.Get("/del/{id:string}", UserGroupDeleteById)
		routes.Get("/getById/{id:string}", UserGroupGetById)
		routes.Post("/list/{pageNum:int}/{pageSize:int}", UserGroupGetList)
	}

}
