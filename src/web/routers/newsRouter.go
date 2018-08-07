package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func NewsRouter(app *iris.Application) {
	routes := app.Party("/v1/news", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", NewsCreate)
		routes.Post("/update", NewsUpdate)
		routes.Get("/del/{id:string}", NewsDeleteById)
		routes.Get("/getById/{id:string}", NewsGetById)
		routes.Post("/list/{pageNum:int}/{pageSize:int}", NewsGetList)
	}

}
