package routers

import (
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func NewsReviewRouter(app *iris.Application) {
	routes := app.Party("/v1/new/review", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", NewsReviewCreate)
		routes.Get("/del/{id:string}", NewsReviewDeleteById)
		routes.Get("/getById/{id:string}", NewsReviewGetById)
		routes.Post("/list/{pageNum:int}/{pageSize:int}", NewsReviewGetList)
	}

}
