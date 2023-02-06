package weixin

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func GetUserOpenidAndSessionKey(code string) (string, string, string) {
	// 请求微信服务器
	OpenIdUrl := utils.FormatString(GetOpenIdUrl, map[string]interface{}{"appid": APPID, "secret": AppSecret, "code": code})
	CodeResp := CodeResult{}
	err := requests.Get(OpenIdUrl, &CodeResp)
	if err != nil {
		msg := "请求微信服务器失败," + err.Error()
		logs.Error(msg)
		return "", "", msg
	}
	return CodeResp.Openid, CodeResp.SessionKey, ""
}
func DecryptPhone(encryptedData string, ivString string, sessionKey string) (phone string, err error) {
	// 解密手机号
	ciphertext, _ := base64.StdEncoding.DecodeString(encryptedData)
	iv, _ := base64.StdEncoding.DecodeString(ivString)
	key, _ := base64.StdEncoding.DecodeString(sessionKey)
	plaintext := make([]byte, len(ciphertext))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	expectedResponseBody := map[string]interface{}{}
	result := PKCS7UnPadding(plaintext)
	json.Unmarshal(result, &expectedResponseBody)
	phone = expectedResponseBody["purePhoneNumber"].(string)
	return phone, err
}
