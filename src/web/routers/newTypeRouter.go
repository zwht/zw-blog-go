package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func NewTypeRouter(app *iris.Application) {
	routes := app.Party("/v1/newType", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", NewTypeCreate)
		routes.Post("/update", NewTypeUpdate)
		routes.Get("/del/{id:string}", NewTypeDeleteById)
		routes.Get("/getById/{id:string}", NewTypeGetById)
		routes.Post("/list/{pageNum:int}/{pageSize:int}", NewTypeGetList)
	}

}
