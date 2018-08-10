package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func VpnRelationRouter(app *iris.Application) {
	routes := app.Party("/v1/vps", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", VpnRelationCreate)
		routes.Post("/update", VpnRelationUpdate)
		routes.Get("/del/{id:string}", VpnRelationDeleteById)
		routes.Get("/getById/{id:string}", VpnRelationGetById)
		routes.Post("/list/{pageNum:int}/{pageSize:int}", VpnRelationGetList)
	}

}
