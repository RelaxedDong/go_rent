package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/astaxie/beego"
	"rent_backend/consts"
	"rent_backend/controllers"
	"rent_backend/controllers/house/manager/view_manager"
	"rent_backend/utils"
	"time"
)

type Controller struct {
	controllers.BaseController
}

var MaxLength int = 200 * 1024 * 1024
var ExpireTimeStr string = "10m"

func (request *Controller) GetOssSign() {
	OssSecretKey := beego.AppConfig.String("oss_secret_key") //您的 AccesskeySecret
	currentTime := time.Now()
	// 10分钟过期
	m, _ := time.ParseDuration(ExpireTimeStr)
	expireTime := currentTime.Add(m)
	expireFormat := expireTime.UTC().Format(time.RFC3339)
	policyConf := `{"expiration": "{expireFormat}", "conditions": [["content-length-range", 0, {maxLength}]]}`
	pol := utils.FormatString(policyConf, map[string]interface{}{
		"expireFormat": expireFormat,
		"maxLength":    MaxLength,
	})
	policy := base64.StdEncoding.EncodeToString([]byte(pol))
	key := []byte(OssSecretKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(policy))
	//进行Base64编码
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	request.RestFulSuccess(map[string]interface{}{
		"sign":        signature,
		"policy":      policy,
		"accessKeyId": beego.AppConfig.String("oss_access_key"),
		"upload_path": beego.AppConfig.String("oss_upload_path"),
		"region_host": beego.AppConfig.String("oss_region_host"),
	}, "")
}

func (request *Controller) Selects() {
	apartmentList := utils.MapToNameValueList(consts.ApartMentTypeMap, true, []string{"0"})
	houseTypeList := utils.MapToNameValueList(consts.RentTypeMap, true, []string{})
	request.RestFulSuccess(map[string]interface{}{
		"facility_list": consts.FacilityMap,
		"apartment":     apartmentList,
		"subway":        []string{},
		"house_type":    houseTypeList,
		"filter_conf":   view_manager.BuildSearchFilters(apartmentList, houseTypeList),
	}, "")
}
