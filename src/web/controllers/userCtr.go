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
	"fmt"
	"github.com/kataras/iris/sessions"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func UserRegister(ctx *Context) {
	captcha := ctx.URLParam("captcha")
	result := Result{}
	var user User
	ctx.ReadJSON(&user)

	s := GetSess(5).Start(ctx)
	capcha := s.GetString(user.Email)
	if capcha != "" {
		a := strings.Split(capcha, "#%#")
		if captcha == a[0] {
			_, err1 := UserSelectByName(user.LoginName)
			_, err2 := UserSelectByPhone(user.Phone)
			_, err3 := UserSelectByEmail(user.Email)
			user.Roles = "1004"
			user.ParentId = "vpn"
			user.Name = user.LoginName
			if err1 == nil {
				result.Code = 402
				result.Msg = "昵称已被使用"
			} else if err2 == nil {
				result.Code = 402
				result.Msg = "电话号码已经被使用"
			} else if err3 == nil {
				result.Code = 402
				result.Msg = "邮箱已经被使用"
			} else {
				// md5加密密码
				h := md5.New()
				h.Write([]byte(user.Password))
				user.Password = hex.EncodeToString(h.Sum(nil))
				err := UserInsert(user)
				if err != nil {
					result.Code = 0
					msg := err.Error()
					if msg == "pq: value too long for type character varying(11)" {
						result.Msg = "电话号码错误"
					} else {
						result.Msg = err.Error()
					}

				} else {
					result.Code = 200
					result.Msg = "成功保存用户信息"
				}
			}
		} else {
			result.Code = 402
			result.Msg = "验证码错误"
		}
	} else {
		result.Code = 402
		result.Msg = "验证码未发送"
	}
	ctx.JSON(result)
}
func UserCreate(ctx *Context) {
	result := Result{}
	var user User
	ctx.ReadJSON(&user)
	_, err1 := UserSelectByName(user.LoginName)
	_, err2 := UserSelectByPhone(user.Phone)
	_, err3 := UserSelectByEmail(user.Email)

	if err1 == nil {
		result.Code = 402
		result.Msg = "登录名重复"
		ctx.JSON(result)
	} else if err2 == nil {
		result.Code = 402
		result.Msg = "电话号码已经被使用"
		ctx.JSON(result)
	} else if err3 == nil {
		result.Code = 402
		result.Msg = "邮箱已经被使用"
		ctx.JSON(result)
	} else {
		// md5加密密码
		h := md5.New()
		h.Write([]byte(user.Password))
		user.Password = hex.EncodeToString(h.Sum(nil))
		err := UserInsert(user)

		if err != nil {
			result.Code = 0
			msg := err.Error()
			if msg == "pq: value too long for type character varying(11)" {
				result.Msg = "电话号码错误"
			} else {
				result.Msg = err.Error()
			}

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
	user2, err2 := UserSelectByPhone(user.Phone)
	user3, err3 := UserSelectByEmail(user.Email)

	if err1 == nil && user1.ID != user.ID {
		result.Code = 402
		result.Msg = "登录名重复"
		ctx.JSON(result)
	} else if err2 == nil && user2.ID != user.ID {
		result.Code = 402
		result.Msg = "电话号码已经被使用"
		ctx.JSON(result)
	} else if err3 == nil && user3.ID != user.ID {
		result.Code = 402
		result.Msg = "邮箱已经被使用"
		ctx.JSON(result)
	} else {
		err := UserUpdate(user)
		if err != nil {
			result.Code = 0
			msg := err.Error()
			if msg == "pq: value too long for type character varying(11)" {
				result.Msg = "电话号码错误"
			} else {
				result.Msg = err.Error()
			}
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
		result.Msg = "登录名或密码错误"
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

//邮件发送验证码
func UserCaptchaEmailCtr(ctx *Context) {
	//"97212287@qq.com"
	// 把token保存到redis
	//s := Sess.Start(ctx)
	//s.Set("token", tokenString)
	//fmt.Printf(s.GetString("token"))
	email := ctx.URLParam("email")
	s := GetSess(5).Start(ctx)
	result := Result{}

	capcha := s.GetString(email)
	if capcha != "" {
		a := strings.Split(capcha, "#%#")
		//a[0]=code a[1]=过期时间
		overTime, _ := strconv.ParseInt(a[1], 10, 64)
		if overTime < time.Now().UnixNano()/1e6 {
			vcode, erra := sendCode(s, email, a[0])
			if erra == nil {
				result.Msg = "已经发送" + vcode
				result.Code = 200
			} else {
				result.Msg = erra.Error()
				result.Code = 430
			}

		} else {
			result.Msg = "验证吗已经发送，在过" + strconv.FormatInt((overTime-time.Now().UnixNano()/1e6)/1000, 10) + "s在发送"
			result.Code = 405
		}
	} else {

		vcode, errb := sendCode(s, email, "")
		if errb == nil {
			result.Msg = "已经发送" + vcode
			result.Code = 200
		} else {
			result.Msg = errb.Error()
			result.Code = 430
		}
	}
	ctx.JSON(result)
}

func sendCode(s *sessions.Session, email string, code string) (vcode string, err error) {
	if code != "" {
		vcode = code
	} else {
		//获取6位数字
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		vcode = fmt.Sprintf("%06v", rnd.Int31n(1000000))
	}
	// 创建redis保存数据，需要过期时间
	tim := time.Now().UnixNano()/1e6 + 1000*60
	s.Set(email, vcode+"#%#"+strconv.FormatInt(tim, 10))
	err = sendEmail(email, vcode)
	return
}
func sendEmail(email string, code string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "1512763623@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "vpn用户注册验证码")
	m.SetBody("text/html", "<div style='width: 600px; margin: 50px auto'><p>您好：</p><h3 style=' line-height: 50px;'>您的vpn注册验证码已经成功发送</h3><p>如果您未做过此申请并认为有人未经授权使用了您的邮件，您不必要把验证码透露给任何人。如果若希望注册vpn账号，然后前往<a style='color: #08c;' href='http://120.79.171.251:9876/'>vpn翻墙神器</a> </p><p style=' line-height: 40px'>您的验证码：<b style='color: #099'>"+code+"</b></p><p style='line-height: 50px;'>此致</p><p>神器vpn 支持</p></div>")
	d := gomail.NewDialer("smtp.qq.com", 465, "1512763623@qq.com", "btdvleiwkbjthjdg")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {

	}
	return
}
