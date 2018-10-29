package routers

import (
	. "../../tools/http"
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func NewsRouter(app *iris.Application) {
	routes := app.Party("/v1/news", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", Permission(NewsCreate, ""))
		routes.Post("/update", Permission(NewsUpdate, ""))
		routes.Get("/del/{id:string}", Permission(NewsDeleteById, ""))
		routes.Get("/getById/{id:string}", Permission(NewsGetById, ""))
		routes.Post("/list/{pageNum:int}/{pageSize:int}", Permission(NewsGetList, ""))
	}

}
