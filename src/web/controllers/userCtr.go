package controllers

import (
	. "../../config"
	. "../../datamodels"
	. "../../services"
	"encoding/json"
	"fmt"
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
func UpdateUser(ctx iris.Context) {
	var user User
	ctx.ReadJSON(&user)

	err := user.Update()

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

		result.Code = 200
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
		userJson, _ := json.Marshal(user)
		// var empList User
		// err3 := json.Unmarshal([]byte(userJson), &empList)
		tokenString, err1 := SetJwt(userJson)
		tokenStr33 := GetJwt(tokenString)
		var empList User
		err3 := json.Unmarshal(tokenStr33, &empList)
		fmt.Printf(err3.Error())

		if err1 != nil {
			result.Code = 0
			result.Msg = err1.Error()
		} else {
			result.Code = 200
			result.Msg = "登录成功"
			// 把token保存到redis
			s := Sess.Start(ctx)
			s.Set("token", tokenString)
			fmt.Printf(s.GetString("token"))

			var userVo map[string]string /*创建集合 */
			userVo = make(map[string]string)
			userVo["name"] = user.Name
			userVo["id"] = user.ID
			userVo["loginName"] = user.LoginName
			userVo["phone"] = user.Phone
			userVo["email"] = user.Email
			userVo["token"] = tokenString
			result.Data = userVo
		}

	}
	ctx.JSON(result)
}
