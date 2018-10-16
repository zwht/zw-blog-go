package routersBlog

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func RoutersBlogInit(app *iris.Application) {
	user := mvc.New(app.Party("/"))
	user.Handle(new(HelloController))
}
