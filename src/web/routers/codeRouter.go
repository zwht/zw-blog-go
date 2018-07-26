package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func CodeRouter(app *iris.Application) {
	routes := app.Party("/v1/code", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", CodeCreate)
		routes.Post("/update", CodeUpdate)
		routes.Get("/del/{id:string}", CodeDeleteById)
		routes.Get("/getById/{id:string}", CodeGetById)
		routes.Post("/list/{pageSize:int}/{pageNum:int}", CodeGetList)
	}

}
