package house

import (
	"rent_backend/consts"
	"rent_backend/controllers"
	"rent_backend/controllers/house/manager/db_manager"
	"rent_backend/controllers/house/manager/view_manager"
	"rent_backend/utils"
	"strconv"
)

type Controller struct {
	controllers.BaseController
}

func (request *Controller) CityListConf() {
	var CityList []map[string]string
	var CityImageUrl = "https://img.donghao.club/conf/city/{city}.png"
	for city, province := range consts.CityMap {
		CityList = append(CityList, map[string]string{
			"province": province,
			"city":     city + "市",
			"url":      utils.FormatString(CityImageUrl, map[string]interface{}{"city": city}),
		})
	}
	request.RestFulSuccess(map[string]interface{}{"city_list": CityList}, "")
}

func (request *Controller) HouseIndex() {
	city := request.Input().Get("city")
	houses := db_manager.GetHouseByQuery(city, "-is_delicate", 10, 0)
	request.RestFulSuccess(map[string]interface{}{"house": view_manager.GetHouseListInfo(houses)}, "")
}

func (request *Controller) HouseDetail() {
	HouseIdString := request.Ctx.Input.Param(":house_id")
	HouseId, _ := strconv.ParseInt(HouseIdString, 10, 64)
	HouseInfo, err := db_manager.GetHouseById(HouseId)
	if err != nil {
		request.RestFulParamsError("房源不存在...", consts.STATUS_CODE_404)
	}
	Info := view_manager.BuildHouseInfo(HouseInfo)
	request.RestFulSuccess(map[string]interface{}{
		"house":           Info,
		"facilities_list": consts.FacilityMap,
	}, "")
}
