package tools

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/satori/go.uuid"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SendCode(target string, codeType string, code string) (vcode string, err error) {
	if code != "" {
		vcode = code
	} else {
		//获取6位数字
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		vcode = fmt.Sprintf("%06v", rnd.Int31n(1000000))
	}
	if codeType == "phone" {
		err = sendPhone(target, vcode)
	}
	if codeType == "email" {
		err = sendEmail(target, vcode)
	}
	return
}

// 手机发送验证码
const sortQueryString_fmt string = "AccessKeyId=%s" +
	"&Action=SendSms" +
	"&Format=JSON" +
	"&OutId=123" +
	"&PhoneNumbers=%s" +
	"&RegionId=cn-hangzhou" +
	"&SignName=%s" +
	"&SignatureMethod=HMAC-SHA1" +
	"&SignatureNonce=%s" +
	"&SignatureVersion=1.0" +
	"&TemplateCode=%s" +
	"&TemplateParam=%s" +
	"&Timestamp=%s" +
	"&Version=2017-05-25"

func encode_local(encode_str string) string {
	urlencode := url.QueryEscape(encode_str)
	urlencode = strings.Replace(urlencode, "+", "%%20", -1)
	urlencode = strings.Replace(urlencode, "*", "%2A", -1)
	urlencode = strings.Replace(urlencode, "%%7E", "~", -1)
	urlencode = strings.Replace(urlencode, "/", "%%2F", -1)
	return urlencode
}
func sendPhone(phone string, code string) (err error) {
	const token string = "fvDzl3yu36pYjzMrs40PttPPag6stx&" // 阿里云 accessSecret 注意这个地方要添加一个 &
	AccessKeyId := "LTAIKcOv9Dw0uZhP"                      // 自己的阿里云 accessKeyID
	PhoneNumbers := phone                                  // 发送目标的手机号
	SignName := url.QueryEscape("赵伟")
	SignatureNonce, _ := uuid.NewV4()
	TemplateCode := "SMS_142384430"
	TemplateParam := url.QueryEscape("{\"code\":\"" + code + "\"}")
	Timestamp := url.QueryEscape(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	sortQueryString := fmt.Sprintf(sortQueryString_fmt,
		AccessKeyId,
		PhoneNumbers,
		SignName,
		SignatureNonce,
		TemplateCode,
		TemplateParam,
		Timestamp,
	)
	urlencode := encode_local(sortQueryString)
	sign_str := fmt.Sprintf("GET&%%2F&%s", urlencode)
	key := []byte(token)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(sign_str))
	signture := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	signture = encode_local(signture)
	fmt.Println(signture)
	fmt.Printf("http://dysmsapi.aliyuncs.com/?Signature=%s&%s\n", signture, sortQueryString)
	st := "http://dysmsapi.aliyuncs.com/?Signature=" + signture + "&" + sortQueryString
	resp, err := http.Get(st)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	return
}

// 邮箱发送验证码
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
