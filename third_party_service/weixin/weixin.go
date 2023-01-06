package weixin

import (
	"github.com/astaxie/beego"
	"rent_backend/utils"
	"rent_backend/utils/requests"
)

var GetOpenIdUrl = "https://api.weixin.qq.com/sns/jscode2session?appid={appid}&secret={secret}&js_code={code}&grant_type=authorization_code"
var APPID = beego.AppConfig.String("APPID")
var AppSecret = beego.AppConfig.String("APP_SECRET")

type CodeResult struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

func GetUserOpenidAndSessionKey(code string) (openid string, errMsg string) {
	// 请求微信服务器
	OpenIdUrl := utils.FormatString(GetOpenIdUrl, map[string]interface{}{"appid": APPID, "secret": AppSecret, "code": code})
	CodeResp := CodeResult{}
	err := requests.Get(OpenIdUrl, &CodeResp)
	if err != nil {
		return "", "请求微信服务器失败," + err.Error()
	}
	return CodeResp.Openid, ""
}
