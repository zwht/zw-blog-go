package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func VpsRouter(app *iris.Application) {
	routes := app.Party("/v1/vps", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", VpsCreate)
		routes.Post("/update", VpsUpdate)
		routes.Get("/del/{id:string}", VpsDeleteById)
		routes.Get("/getById/{id:string}", VpsGetById)
		routes.Post("/list/{pageNum:int}/{pageSize:int}", VpsGetList)
	}

}
