package main

import (
	. "./tools/http"
	. "./web/routers"
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	app := iris.New()
	//api文档配置
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "restfull api Iris",
		DocPath:  "web/views/api.html",
		BaseUrls: map[string]string{"Production": "go restfull base", "Staging": "ris"},
	})
	app.Use(irisyaag.New()) // <- IMPORTANT, register the middleware.

	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	//静态文件配置
	app.StaticWeb("/", "./web/views/dist")
	//html模板配置
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//项目入口
	app.Get("/", func(ctx iris.Context) {
		ctx.View("dist/index.html")
	})
	//qpi入口
	app.Get("/api", func(ctx iris.Context) {
		ctx.View("api.html")
	})

	//RedisInit()
	RouterInit(app)
	//http拦截器中间件
	app.UseGlobal(HttpInterceptor)

	app.Run(iris.Addr(":8888"),
		iris.WithCharset("UTF-8"),
		iris.WithoutServerError(iris.ErrServerClosed))
}
