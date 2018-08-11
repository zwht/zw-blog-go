package routers

import (
	. "../../tools/http"
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func VpnRouter(app *iris.Application) {
	routes := app.Party("/v1/vpn", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", Permission(VpnCreate, "1001"))
		routes.Post("/update", Permission(VpnUpdate, "1001"))
		//routes.Get("/updateState", VpnUpdateStateCtr)
		routes.Get("/del/{id:string}", Permission(VpnDeleteById, "1001"))
		routes.Get("/getById/{id:string}", Permission(VpnGetById, "1001"))
		routes.Post("/list/{pageNum:int}/{pageSize:int}", Permission(VpnGetList, "1001"))
	}

}
