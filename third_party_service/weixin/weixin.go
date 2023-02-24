package weixin

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"rent_backend/utils/requests"
	"strconv"
)

var APPID = beego.AppConfig.String("APPID")
var AppSecret = beego.AppConfig.String("APP_SECRET")

const (
	qrcodeMiniPath = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s"
	openIdUrl      = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	accessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

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
	url := fmt.Sprintf(openIdUrl, APPID, AppSecret, code)
	CodeResp := CodeResult{}
	err := requests.Get(url, &CodeResp)
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

func GetAccessToken() (string, error) {
	//生成一个基础access_token
	url := fmt.Sprintf(accessTokenUrl, APPID, AppSecret)
	resp := map[string]interface{}{}
	err := requests.Get(url, &resp)
	if err != nil {
		logs.Error("获取AccessToken错误" + err.Error())
		return "", err
	}
	return resp["access_token"].(string), nil
}

func GetPathImgByHouseId(houseId int64) (string, error) {
	// 从微信服务器获取小程序的二维码路径
	accessToken, err := GetAccessToken()
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf(qrcodeMiniPath, accessToken)
	data, err := requests.JsonPostGetBody(url, map[string]string{
		"scene": "house=" + strconv.FormatInt(houseId, 10),
		"page":  "pages/detail/detail",
	})
	if err != nil {
		logs.Error("获取小程序path出错" + err.Error())
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), err
}
