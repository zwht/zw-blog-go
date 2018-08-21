package controllers

import (
	. "../../datamodels"
	. "../../models"
	. "../../services"
	. "../../tools"
	. "../../tools/http"
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
	"time"
)

//使用邮件注册
func UserRegisterEmailCtr(ctx *Context) {
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
			_, err3 := UserSelectByEmail(user.Email)
			user.Roles = "1004"
			user.ParentId = "vpn"
			user.Name = user.LoginName
			if err1 == nil {
				result.Code = 402
				result.Msg = "昵称已被使用"
			} else if err3 == nil {
				result.Code = 402
				result.Msg = "邮箱已经被使用"
			} else {
				// md5加密密码
				h := md5.New()
				h.Write([]byte(user.Password))
				user.Password = hex.EncodeToString(h.Sum(nil))
				user.ID = uuid.Must(uuid.NewV4()).String()
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
			vcode, erra := SendCode(email, "phone", a[0])
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
		vcode, errb := SendCode(email, "phone", "")
		// 创建redis保存数据，需要过期时间
		tim := time.Now().UnixNano()/1e6 + 1000*60
		s.Set(email, vcode+"#%#"+strconv.FormatInt(tim, 10))

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

//使用手机号注册
func UserRegisterPhoneCtr(ctx *Context) {
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
			user.Roles = "1004"
			user.ParentId = "vpn"
			user.Name = user.LoginName
			if err1 == nil {
				result.Code = 402
				result.Msg = "昵称已被使用"
			} else if err2 == nil {
				result.Code = 402
				result.Msg = "电话号码已经被使用"
			} else {
				// md5加密密码
				h := md5.New()
				h.Write([]byte(user.Password))
				user.Password = hex.EncodeToString(h.Sum(nil))
				user.ID = uuid.Must(uuid.NewV4()).String()
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

//手机号发送验证码
func UserCaptchaPhoneCtr(ctx *Context) {
	result := Result{}
	ctx.JSON(result)
}
