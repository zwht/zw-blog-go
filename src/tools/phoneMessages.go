package tools

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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
func SendPhoneMessages() {
	const token string = "fvDzl3yu36pYjzMrs40PttPPag6stx&" // 阿里云 accessSecret 注意这个地方要添加一个 &
	AccessKeyId := "LTAIKcOv9Dw0uZhP"                      // 自己的阿里云 accessKeyID
	PhoneNumbers := "13540416116"                          // 发送目标的手机号
	SignName := url.QueryEscape("赵伟")
	SignatureNonce, _ := uuid.NewV4()
	TemplateCode := "SMS_142384430"
	TemplateParam := url.QueryEscape("{\"code\":\"888888\"}")
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
}
