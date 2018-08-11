package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func VpnRouter(app *iris.Application) {
	routes := app.Party("/v1/vpn", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", VpnCreate)
		routes.Post("/update", VpnUpdate)
		routes.Get("/del/{id:string}", VpnDeleteById)
		routes.Get("/getById/{id:string}", VpnGetById)
		routes.Post("/list/{pageNum:int}/{pageSize:int}", VpnGetList)
	}

}
