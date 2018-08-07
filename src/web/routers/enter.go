package routers

import (
	"github.com/kataras/iris"
)

func RouterInit(app *iris.Application) {
	UserRouter(app)
	UserGroupRouter(app)
	CodeRouter(app)
	NewsTypeRouter(app)
	NewsReviewRouter(app)
	NewsRouter(app)
}
