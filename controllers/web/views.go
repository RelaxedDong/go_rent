package web

import (
	"rent_backend/consts"
	"rent_backend/controllers"
	"rent_backend/controllers/house/manager/db_manager"
	"rent_backend/controllers/house/manager/view_manager"
	"rent_backend/third_party_service/weixin"
	"rent_backend/utils/datetime"
	"strconv"
	"time"
)

type Controller struct {
	controllers.BaseController
}

func (request *Controller) HouseDetail() {
	// 这里不用处理错误，路由会直接拦截报Not Found
	HouseIdString := request.Ctx.Input.Param(":house_id")
	HouseId, _ := strconv.ParseInt(HouseIdString, 10, 64)
	HouseInfo, _ := db_manager.GetHouseById(HouseId)
	HouseData := view_manager.BuildHouseInfo(HouseInfo).(map[string]interface{})
	request.Data["house"] = HouseData
	request.Data["default_image"] = consts.DEFAULT_NONE_IMAGE
	request.Data["mini_img"], _ = weixin.GetPathImgByHouseId(HouseId)
	request.TplName = "detail.html"
}

func (request *Controller) Index() {
	request.TplName = "index.html"
}

func (request *Controller) AskRentIndex() {
	request.TplName = "ask_rent.html"
}

func (request *Controller) RentStatistics() {
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	request.Data["statistics"] = db_manager.GetUserStatistics(WxUser.Id)
	request.Data["Now"] = datetime.TimeToStr(time.Now())
	request.TplName = "house_statistics.html"
}
