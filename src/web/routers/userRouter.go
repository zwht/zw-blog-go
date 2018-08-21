package routers

import (
	. "../../tools/http"
	. "../controllers"
	"github.com/kataras/iris"
)

//user model
func UserRouter(app *iris.Application) {
	routes := app.Party("/v1/user", func(ctx iris.Context) {

		ctx.Next()
	})
	{
		routes.Post("/add", Permission(UserCreate, "1001"))
		routes.Post("/update", Permission(UserUpdateCtr, ""))
		routes.Get("/updateState", Permission(UserUpdateStateCtr, "1001"))
		routes.Get("/updatePassword", Permission(UserUpdatePassowordCtr, ""))
		routes.Get("/del/{id:string}", Permission(UserDeleteById, "1001"))
		routes.Get("/getById/{id:string}", Permission(UserGetById, ""))
		routes.Post("/list/{pageNum:int}/{pageSize:int}", Permission(UserGetList, "1001"))

		routes.Post("/login", Permission(Login, ""))
		routes.Get("/captcha/email", Permission(UserCaptchaEmailCtr, ""))
		routes.Post("/register/email", Permission(UserRegisterEmailCtr, ""))
		routes.Get("/captcha/phone", Permission(UserCaptchaPhoneCtr, ""))
		routes.Post("/register/phone", Permission(UserRegisterPhoneCtr, ""))
	}
}
