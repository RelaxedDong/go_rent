package house_api

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"rent_backend/consts"
	"rent_backend/controllers"
	UserDbManager "rent_backend/controllers/account/manager/db_manager"
	houseform "rent_backend/controllers/house/form"
	"rent_backend/controllers/house/manager/db_manager"
	"rent_backend/controllers/house/manager/view_manager"
	"rent_backend/utils"
	"rent_backend/utils/email"
	"strconv"
)

type Controller struct {
	controllers.BaseController
}

func (request *Controller) CityListConf() {
	var CityList []map[string]string
	var CityImageUrl = "https://img.donghao.club/conf/city/{city}.png"
	for _, city := range consts.CityMapSorted {
		province := consts.CityMap[city]
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
	start, _ := request.GetStartEndByPage(consts.DefaultPageSize)
	houses := db_manager.GetHouseByQuery(city, "", "", []string{}, []string{}, []string{}, []uint64{},
		"-is_delicate", start, consts.DefaultPageSize)
	request.RestFulSuccess(map[string]interface{}{"houses": view_manager.GetHouseListInfo(houses)}, "")
}

func (request *Controller) HouseDetail() {
	isLogin, WxUser := request.GetWxUser()
	// 这里不用处理错误，路由会直接拦截报Not Found
	HouseIdString := request.Ctx.Input.Param(":house_id")
	HouseId, _ := strconv.ParseInt(HouseIdString, 10, 64)
	HouseInfo, err := db_manager.GetHouseById(HouseId)
	if err != nil {
		request.RestFulParamsError(consts.ErrorMsgHouseNotExists, consts.STATUS_CODE_404)
	}
	Info := view_manager.BuildHouseInfo(HouseInfo)
	result := map[string]interface{}{
		"house":           Info,
		"scroll_texts":    []string{},
		"facilities_list": consts.FacilityMap,
	}
	if isLogin {
		result["is_collect"] = UserDbManager.IsUserCollectHouse(HouseId, WxUser.Id)
		go UserDbManager.GetOrCreateUserHistory(HouseInfo, WxUser)
	}
	request.RestFulSuccess(result, "")
}

func (request *Controller) BannerList() {
	city := request.Input().Get("city")
	pagesConfig, _ := db_manager.GetBannerByQuery(city, []string{"banner", "icon"})
	banners, icons := view_manager.GetHomePageConfig(pagesConfig)
	request.RestFulSuccess(map[string]interface{}{
		"banners": banners,
		"icons":   icons,
	}, "")
}

func (request *Controller) SearchHouse() {
	city := request.Input().Get("city")
	//page := request.Input().Get("page")
	start, _ := request.GetStartEndByPage(consts.DefaultPageSize)
	title := request.Input().Get("title")
	filterConfStr := request.Input().Get("filter_conf")
	FilterConf := map[string]interface{}{}
	json.Unmarshal([]byte(filterConfStr), &FilterConf)
	filter := view_manager.GetHouseFilterBySearch(FilterConf)
	// 首页传递过来的参数
	SearchHouseType := request.GetString("house_type")
	if SearchHouseType != "" {
		filter["houseTypes"] = append(filter["houseTypes"].([]string), SearchHouseType)
	}
	houses := db_manager.GetHouseByQuery(city, title, filter["region"].(string),
		filter["houseTypes"].([]string),
		filter["apartments"].([]string),
		filter["facilitiesList"].([]string),
		[]uint64{filter["startPrice"].(uint64), filter["endPrice"].(uint64)},
		filter["sort_by"].(string),
		start,
		consts.DefaultPageSize)
	request.RestFulSuccess(map[string]interface{}{"houses": view_manager.GetHouseListInfo(houses)}, "")
}

func (request *Controller) NearbyHouses() {
	city := request.Input().Get("city")
	locationConf := request.Input().Get("location_conf")
	excludeId := request.Input().Get("exclude_id")
	//start, _ := request.GetStartEndByPage(consts.DefaultPageSize)
	LocationData := map[string]interface{}{}
	json.Unmarshal([]byte(locationConf), &LocationData)
	lng := LocationData["longitude"].(float64)
	lat := LocationData["latitude"].(float64)
	houses := db_manager.GetNearByHouses(city, lng, lat, 5)
	excludeIds := []int64{}
	if excludeId != "" {
		houseId, _ := strconv.Atoi(excludeId)
		excludeIds = append(excludeIds, int64(houseId))
	}
	request.RestFulSuccess(map[string]interface{}{"houses": view_manager.GetHouseListInfo(houses, excludeIds)}, "")
}

func (request *Controller) HouseAdd() {
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	var req houseform.HouseAddForm
	request.RequestJsonFormat(&req)
	subject := "房源审核提醒"
	if req.HouseId == 0 {
		err := db_manager.CreateHouse(req, WxUser)
		if err != nil {
			request.RestFulParamsError("创建房源失败: " + err.Error())
		}
	} else {
		HouseInfo, err := db_manager.GetHouseByIdNoPublisher(req.HouseId)
		if err != nil || HouseInfo.Publisher.Id != WxUser.Id {
			request.RestFulParamsError(consts.ErrorMsgHouseNotExists)
		}
		err = db_manager.UpdateHouse(req.HouseId, req, WxUser)
		if err != nil {
			request.RestFulParamsError("房源更新失败: " + err.Error())
		}
		// 编辑房源
		subject = "房源更新审核提醒"
	}
	// 发送审核邮件
	message := fmt.Sprintf(consts.HouseCheckMessage, WxUser.NickName)
	go email.SendMail(beego.AppConfig.String("email_user"), subject, message)
	request.RestFulSuccess(map[string]interface{}{}, "")
}

func (request *Controller) HouseAddCheck() {
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	if WxUser.Phone == "" || WxUser.Wechat == "" {
		request.RestFulSuccess(map[string]interface{}{
			"can_publish": false,
		},
			"为了方便租客联系，请先绑定信息~")
	}
	request.RestFulSuccess(map[string]interface{}{"can_publish": true}, "")
}

func (request *Controller) HouseDelete() {
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	postData := request.GetPostBody()
	houseId, ok := postData["houseid"]
	if !ok {
		request.RestFulParamsError(consts.ErrorMsgHouseNotExists)
	}
	house, err := db_manager.GetHouseByIdNoPublisher(int64(houseId.(float64)))
	if err != nil || house.Publisher.Id != WxUser.Id {
		request.RestFulParamsError(consts.ErrorMsgHouseNotExists)
	}
	err = db_manager.DeleteHouse(house)
	if err != nil {
		request.RestFulParamsError("删除失败:" + err.Error())
	}
	request.RestFulSuccess(map[string]interface{}{}, "")
}
