package routers

import (
	. "../../tools/http"
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func CodeRouter(app *iris.Application) {
	routes := app.Party("/v1/code", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", Permission(CodeCreate, "1001"))
		routes.Post("/update", Permission(CodeUpdate, "1001"))
		routes.Get("/del/{id:string}", Permission(CodeDeleteById, "1001"))
		routes.Get("/getById/{id:string}", Permission(CodeGetById, "1001"))
		routes.Post("/list/{pageNum:int}/{pageSize:int}", Permission(CodeGetList, ""))
	}

}
