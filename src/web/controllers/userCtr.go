package controllers

import (
	. "../../config"
	. "../../datamodels"
	. "../../models"
	. "../../services"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/kataras/iris"
	"strconv"
)

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
func UserCreate(ctx iris.Context) {
	result := Result{}
	var user User
	ctx.ReadJSON(&user)
	_, err1 := UserSelectByName(user.LoginName)
	if err1 == nil {
		result.Code = 402
		result.Msg = "登录名重复"
		ctx.JSON(result)
	} else {
		// md5加密密码
		h := md5.New()
		h.Write([]byte(user.Password))
		user.Password = hex.EncodeToString(h.Sum(nil))
		err := UserInsert(user)

		if err != nil {
			result.Code = 0
			result.Msg = err.Error()
		} else {
			result.Code = 200
			result.Msg = "成功保存用户信息"
		}

		ctx.JSON(result)
	}
}
func UserUpdateCtr(ctx iris.Context) {
	result := Result{}
	var user User
	ctx.ReadJSON(&user)

	user1, err1 := UserSelectByName(user.LoginName)
	if err1 == nil && user1.ID != user.ID {
		result.Code = 402
		result.Msg = "登录名重复"
		ctx.JSON(result)
	} else {
		err := UserUpdate(user)
		if err != nil {
			result.Code = 0
			result.Msg = err.Error()
		} else {
			result.Code = 200
			result.Msg = "成功保存用户信息"
		}
		ctx.JSON(result)
	}
}

func UserGetById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	user, err := UserSelect(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功获取用户信息"
		result.Data = user
	}

	ctx.JSON(result)
}

func UserGetList(ctx iris.Context) {

	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var userSearchVo UserSearchVo
	ctx.ReadJSON(&userSearchVo)
	count, _ := UserSelectCount(userSearchVo)
	users, err := UserSelectList(pageSize, pageNum, userSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = users

		result.Code = 200
		result.Msg = "成功获取用户列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func UserDeleteById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := UserDelete(id)

	result := Result{}

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "成功删除用户信息"
	}

	ctx.JSON(result)
}

func Login(ctx iris.Context) {

	var loginVo LoginVo
	ctx.ReadJSON(&loginVo)
	h := md5.New()
	h.Write([]byte(loginVo.Password))
	loginVo.Password = hex.EncodeToString(h.Sum(nil))

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

		if err1 != nil {
			result.Code = 0
			result.Msg = err1.Error()
		} else {
			result.Code = 200
			result.Msg = "登录成功"
			// 把token保存到redis
			//s := Sess.Start(ctx)
			//s.Set("token", tokenString)
			//fmt.Printf(s.GetString("token"))

			var userVo map[string]string /*创建集合 */
			userVo = make(map[string]string)
			userVo["name"] = user.Name
			userVo["id"] = user.ID
			userVo["loginName"] = user.LoginName
			userVo["phone"] = user.Phone
			userVo["email"] = user.Email
			userVo["roles"] = user.Roles
			userVo["token"] = tokenString
			result.Data = userVo
		}

	}
	ctx.JSON(result)
}
