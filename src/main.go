package main

import (
	. "./tools"
	. "./tools/http"
	. "./web/routersBlog"
	. "./web/routersRestFull"
	// "github.com/betacraft/yaag/irisyaag"
	// "github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	app := iris.New()
	//api文档配置
	// yaag.Init(&yaag.Config{
	// 	On:       true,
	// 	DocTitle: "restfull api Iris",
	// 	DocPath:  "web/views/api.html",
	// 	BaseUrls: map[string]string{"Production": "go restfull base", "Staging": "ris"},
	// })
	// app.Use(irisyaag.New())

	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	//静态文件配置
	app.StaticWeb("/img/", "./web/views/img")
	app.StaticWeb("/blog/", "./web/views/blog/assets")
	app.StaticWeb("/", "./web/views/dist")
	//html模板配置
	// app.RegisterView(iris.HTML("./web/views/blog", ".html").
	// 	Layout("layout.html").
	// 	Reload(true))
	app.RegisterView(iris.HTML("./web/views", ".html"))

	// 管理平台入口
	app.Get("/admin", func(ctx iris.Context) {
		ctx.View("dist/index.html")
	})
	//qpi入口
	// app.Get("/api", func(ctx iris.Context) {
	// 	ctx.View("api.html")
	// })

	// redis初始化
	RedisInit()
	// restFull接口路由
	RoutersRestFullInit(app)
	// blog接口路由,项目入口
	RoutersBlogInit(app)
	//http拦截器中间件
	app.UseGlobal(HttpInterceptor)

	app.Run(iris.Addr(":9876"),
		iris.WithCharset("UTF-8"),
		iris.WithoutServerError(iris.ErrServerClosed))
}
