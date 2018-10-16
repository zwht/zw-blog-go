package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func NewsTypeRouter(app *iris.Application) {
	routes := app.Party("/v1/news/type", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", NewsTypeCreate)
		routes.Post("/update", NewsTypeUpdate)
		routes.Get("/del/{id:string}", NewsTypeDeleteById)
		routes.Get("/getById/{id:string}", NewsTypeGetById)
		routes.Post("/list/{pageNum:int}/{pageSize:int}", NewsTypeGetList)
	}

}
