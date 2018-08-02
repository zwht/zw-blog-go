package controllers

import (
	. "../../datamodels"
	. "../../models"
	. "../../services"
	. "../../tools"
	. "../../tools/http"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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
func UserCreate(ctx *Context) {
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
func UserUpdateCtr(ctx *Context) {
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
func UserUpdateStateCtr(ctx *Context) {
	//id := ctx.Params().Get("id")
	st := ctx.URLParam("state")
	id := ctx.URLParam("id")

	state := false
	if st == "true" {
		state = true
	}

	err := UserUpdateState(id, state)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "修改状态成功"
	}
	ctx.JSON(result)
}

func UserUpdatePassowordCtr(ctx *Context) {
	//id := ctx.Params().Get("id")
	oldPassword := ctx.URLParam("oldPassword")
	password := ctx.URLParam("password")
	id := ctx.URLParam("id")
	if id == "" {
		id = ctx.User.ID
	}
	h := md5.New()
	h.Write([]byte(password))
	password = hex.EncodeToString(h.Sum(nil))

	h2 := md5.New()
	h2.Write([]byte(oldPassword))
	oldPassword = hex.EncodeToString(h2.Sum(nil))

	err := UserUpdatePassoword(id, oldPassword, password)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = "密码错误"
	} else {
		result.Code = 200
		result.Msg = "修改成功"
	}
	ctx.JSON(result)
}

func UserGetById(ctx *Context) {
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

func UserGetList(ctx *Context) {

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
		//删除自己数据
		var j = -1
		for i, v := range users {
			if v.ID == ctx.User.ID {
				j = i
			}
		}
		if j != -1 {
			users = append(users[:j], users[j+1:]...)
		}
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

func UserDeleteById(ctx *Context) {
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

func Login(ctx *Context) {

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
			result.Msg = "登录名或密码错误"
		} else {
			if user.State {
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
			} else {
				result.Code = 403
				result.Msg = "账号被禁用"
			}
		}
	}
	ctx.JSON(result)
}
