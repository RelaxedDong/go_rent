package house

import (
	"rent_backend/controllers"
	"rent_backend/utils"
)

type Controller struct {
	controllers.BaseController
}

var CityMap = map[string]string{
	"北京": "北京市",
	"上海": "上海市",
	"广州": "广东省",
	"深圳": "广东省",
	"杭州": "浙江省",
	"成都": "四川省",
	"武汉": "湖北省",
	"长沙": "湖南省",
	"郑州": "河南省",
	"西安": "陕西省",
	"天津": "天津省",
	"厦门": "福建省",
}
var CityImageUrl = "https://img.donghao.club/conf/city/{city}.png"

func (request *Controller) CityListConf() {
	var CityList []map[string]string
	for city, province := range CityMap {
		CityList = append(CityList, map[string]string{
			"province": province,
			"city":     city + "市",
			"url":      utils.FormatString(CityImageUrl, map[string]interface{}{"city": city}),
		})
	}
	request.RestFulSuccess(map[string]interface{}{"city_list": CityList}, "")
}
