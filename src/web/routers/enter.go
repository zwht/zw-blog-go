package routers

import (
	"github.com/kataras/iris"
)

func RouterInit(app *iris.Application) {
	UserRouter(app)
}
