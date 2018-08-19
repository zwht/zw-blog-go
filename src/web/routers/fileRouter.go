package routers

import (
	. "../../tools/http"
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func FileRouter(app *iris.Application) {
	const maxSize = 5 << 20 // 5MB
	routes := app.Party("/v1/file", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/upload", iris.LimitRequestBodySize(maxSize+1<<20), Permission(FileUpload, ""))
		routes.Get("/img", Permission(FileGet, ""))
	}

}
