package main

import (
	. "./web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	app.RegisterView(iris.HTML("./public", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	}) //users
	usersRoutes := app.Party("/users", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		usersRoutes.Post("/", CreateUser)
		usersRoutes.Delete("/{id:int}", DeleteUserById)
		usersRoutes.Get("/{id:int}", GetUserById)
		usersRoutes.Get("/", GetUserList)
	}

	app.Run(iris.Addr(":8080"),
		iris.WithCharset("UTF-8"),
		iris.WithoutServerError(iris.ErrServerClosed))
}
