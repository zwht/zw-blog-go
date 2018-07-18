package hh

import (
	. "../service"
	. "../tool"
	"github.com/kataras/iris"
	"strconv"
)

func CreateUser(ctx iris.Context) {
	var user User
	ctx.ReadJSON(&user)

	err := user.Insert()

	result := Result{}

	if err != nil {
		result.No = 0
		result.Msg = err.Error()
	} else {
		result.No = 1
		result.Msg = "成功保存用户信息"
	}

	ctx.JSON(result)
}

func GetUserById(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	user, err := Select(id)

	result := Result{}

	if err != nil {
		result.No = 0
		result.Msg = err.Error()
	} else {
		result.No = 1
		result.Msg = "成功获取用户信息"
		result.Obj = user
	}

	ctx.JSON(result)
}

func GetUserList(ctx iris.Context) {
	users, err := SelectList()

	result := Result{}

	if err != nil {
		result.No = 0
		result.Msg = err.Error()
	} else {
		result.No = 1
		result.Msg = "成功获取用户列表信息"
		result.Obj = users
	}

	ctx.JSON(result)
}

func DeleteUserById(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	err := Delete(id)

	result := Result{}

	if err != nil {
		result.No = 0
		result.Msg = err.Error()
	} else {
		result.No = 1
		result.Msg = "成功删除用户信息"
	}

	ctx.JSON(result)
}
