package routers

import (
	. "../../tools/http"
	. "../controllers"
	"github.com/kataras/iris"
)

//news_review model
func NewsReviewRouter(app *iris.Application) {
	routes := app.Party("/v1/news_review", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", Permission(NewsReviewCreate, ""))
		routes.Post("/update", Permission(NewsReviewUpdate, "1001"))
		routes.Post("/updateState", Permission(NewsReviewUpdateState, "1001"))
		routes.Get("/del/{id:string}", Permission(NewsReviewDeleteById, "1001"))
		routes.Get("/getById/{id:string}", Permission(NewsReviewGetById, ""))
		routes.Post("/list/{pageNum:int}/{pageSize:int}", Permission(NewsReviewGetList, ""))
	}
}
