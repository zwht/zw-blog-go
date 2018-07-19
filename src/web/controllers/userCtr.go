package controllers

import (
	. "../../datamodels"
	. "../../services"
	"github.com/kataras/iris"
)

func CreateUser(ctx iris.Context) {
	var user User
	ctx.ReadJSON(&user)

	err := user.Insert()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 1
		result.Msg = "成功保存用户信息"
	}

	ctx.JSON(result)
}

func GetUserById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	user, err := Select(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 1
		result.Msg = "成功获取用户信息"
		result.Data = user
	}

	ctx.JSON(result)
}

func GetUserList(ctx iris.Context) {
	users, err := SelectList()

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = 10
		resultPage.PageData = users

		result.Code = 1
		result.Msg = "成功获取用户列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func DeleteUserById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := Delete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 1
		result.Msg = "成功删除用户信息"
	}

	ctx.JSON(result)
}

func Login(ctx iris.Context) {
	var loginVo LoginVo
	ctx.ReadJSON(&loginVo)
	user, err := Logins(loginVo)
	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "登录成功"
		result.Data = user
	}

	ctx.JSON(result)
}