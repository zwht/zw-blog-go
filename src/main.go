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
	//api文档配置//http://localhost:9876/api
	// yaag.Init(&yaag.Config{
	// 	On:       true,
	// 	DocTitle: "restfull api Iris",
	// 	DocPath:  "web/views/api/index.html",
	// 	BaseUrls: map[string]string{"Production": "go restfull base", "Staging": "ris"},
	// })
	// app.Use(irisyaag.New())

	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	//静态文件配置
	app.StaticWeb("/admin/", "./web/views/admin") //管理平台入口
	app.StaticWeb("/api/", "./web/views/api")     //api入口
	app.StaticWeb("/img/", "./web/views/img")
	app.StaticWeb("/vb/", "./web/views/blog/assets")
	//html模板配置
	htmEng := iris.HTML("./web/views", ".html").Reload(true)
	htmEng.AddFunc("RenderHtml", RenderHtml)
	app.RegisterView(htmEng)

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
