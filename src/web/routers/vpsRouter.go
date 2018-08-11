package routers

import (
	. "../../tools/http"
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func VpsRouter(app *iris.Application) {
	routes := app.Party("/v1/vps", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", Permission(VpsCreate, "1001"))
		routes.Post("/update", Permission(VpsUpdate, "1001"))
		routes.Get("/del/{id:string}", Permission(VpsDeleteById, "1001"))
		routes.Get("/getById/{id:string}", Permission(VpsGetById, "1001"))
		routes.Post("/list/{pageNum:int}/{pageSize:int}", Permission(VpsGetList, "1001"))
	}

}
