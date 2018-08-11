package routers

import (
	. "../../tools/http"
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func VpnRelationRouter(app *iris.Application) {
	routes := app.Party("/v1/vpn/relation", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", Permission(VpnRelationCreate, "1001"))
		routes.Post("/update", Permission(VpnRelationUpdate, "1001"))
		routes.Get("/updateState", Permission(VpnRelationUpdateStateCtr, "1001"))
		routes.Get("/del/{id:string}", Permission(VpnRelationDeleteById, "1001"))
		routes.Get("/getById/{id:string}", Permission(VpnRelationGetById, "1001"))
		routes.Post("/list/{pageNum:int}/{pageSize:int}", Permission(VpnRelationGetList, "1001"))
		routes.Get("/list", Permission(VpnRelationGetListByUserId, "1001"))
	}

}
