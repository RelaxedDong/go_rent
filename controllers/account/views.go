package account

import (
	"rent_backend/consts"
	"rent_backend/controllers"
	accountform "rent_backend/controllers/account/form"
	UserDbManager "rent_backend/controllers/account/manager/db_manager"
	"rent_backend/controllers/house/manager/db_manager"
	"rent_backend/controllers/house/manager/view_manager"
	"rent_backend/models"
	_ "rent_backend/models"
	"rent_backend/third_party_service/weixin"
	"rent_backend/utils/jwt"
)

type Controller struct {
	controllers.BaseController
}

func (request *Controller) Login() {

	var req accountform.LoginForm
	data := make(map[string]interface{})
	request.RequestJsonFormat(&req)
	openId, SessionKey, errMsg := weixin.GetUserOpenidAndSessionKey(req.Code)
	if errMsg != "" {
		request.RestFulParamsError(errMsg)
	}
	jwtToken, err := jwt.GenerateToken(openId)
	if err != nil {
		request.RestFulParamsError(err.Error())
	}
	data["token"] = jwtToken
	account, userError := UserDbManager.GetUserByOpenId(openId)
	if userError != nil {
		request.RestFulSuccess(map[string]interface{}{"token": jwtToken}, userError.Error())
	}
	data["user_id"] = account.Id
	UserDbManager.UpdateUserInfo(account, "", "", SessionKey, true)
	// todo: is_superuser, finish_user_info字段舍弃？
	request.RestFulSuccess(data, "")
}

// BindUserInfo 这里是用户点击授权绑定处理，如果当前用户是第一次授权绑定，则创建一个新用户
func (request *Controller) BindUserInfo() {
	WxUser := request.Ctx.Input.GetData("WxUser")
	var UserId int64
	// 不存在就创建一个用户
	if WxUser == nil {
		// 断言值是给定的类型
		OpenId := request.Ctx.Input.GetData("openId").(string)
		var req accountform.UserInfoForm
		req.OpenId = OpenId
		request.RequestJsonFormat(&req)
		_, UserId = UserDbManager.GetOrCreateUser(req)
	} else {
		UserId = WxUser.(models.AccountModel).Id
	}
	request.RestFulSuccess(map[string]interface{}{"user_id": UserId}, "")
}

func (request *Controller) UserInfo() {
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	request.RestFulSuccess(map[string]interface{}{
		"nickname":  WxUser.NickName,
		"avatarUrl": WxUser.AvatarUrl,
		"wechat":    WxUser.Wechat,
		"phone":     WxUser.Phone,
		"gender":    WxUser.Gender,
	}, "")
}

func (request *Controller) BindUserPhone() {
	// todo: 通过微信api直接按钮拿到phone
}

func (request *Controller) EditUserInfo() {
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	var req accountform.EditUserInfoForm
	request.RequestJsonFormat(&req)
	UserDbManager.UpdateUserInfo(WxUser, req.Wechat, req.Phone, "", false)
	request.RestFulSuccess(map[string]interface{}{}, "")
}

func (request *Controller) Operation() {
	// 数据==管理操作
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	var req accountform.OperationForm
	request.RequestJsonFormat(&req)
	if req.OperationType == "collect" {
		HouseInfo, err := db_manager.GetHouseById(req.HouseId)
		if err != nil {
			request.RestFulParamsError(consts.ErrorMsgHouseNotExists, consts.STATUS_CODE_404)
		}
		isNew, RecordId := UserDbManager.GetOrCreateUserFavor(HouseInfo, WxUser)
		if !isNew {
			UserDbManager.DeleteCollectRecordDb(RecordId)
		}
	}
	request.RestFulSuccess(map[string]interface{}{}, "")
}

func (request *Controller) BindPhone() {
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	var req accountform.BindPhoneForm
	request.RequestJsonFormat(&req)
	phone, err := weixin.DecryptPhone(req.EncryptedData, req.Iv, WxUser.SessionKey)
	if err != nil {
		request.RestFulParamsError(err.Error())
	}
	UserDbManager.UpdateUserInfo(WxUser, "", phone, "", false)
	request.RestFulSuccess(map[string]interface{}{"phone": phone}, "")
}

func (request *Controller) Collects() {
	// 获取历史记录与收藏
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	opType := request.GetString("op_type")
	start, _ := request.GetStartEndByPage(consts.DefaultPageSize)
	var houses []models.HouseModel
	switch opType {
	case "collect":
		houses = UserDbManager.GetUserCollects(WxUser.Id, start, consts.DefaultPageSize)
	case "history":
		houses = UserDbManager.GetUserHistoryList(WxUser.Id, start, consts.DefaultPageSize)
	}
	request.RestFulSuccess(map[string]interface{}{"houses": view_manager.GetHouseListInfo(houses)}, "")
}

func (request *Controller) OperationDelete() {
	// 获取历史记录与收藏
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	var req accountform.OperationDeleteForm
	request.RequestJsonFormat(&req)
	var houses []models.HouseModel
	switch req.OperationType {
	case "collect":
		if req.HouseId != 0 {
			UserDbManager.DeleteUserCollectRecord(req.HouseId, WxUser.Id)
		} else {
			UserDbManager.DeleteAllCollectsByUserId(WxUser)
		}
	case "history":
		if req.HouseId != 0 {
			UserDbManager.DeleteHistoryRecordDb(req.HouseId, WxUser.Id)
		} else {
			UserDbManager.DeleteAllHistoryByUserId(WxUser)
		}
	}
	request.RestFulSuccess(map[string]interface{}{"houses": view_manager.GetHouseListInfo(houses)}, "")
}

func (request *Controller) UserPublish() {
	// 获取历史记录与收藏
	request.LoginRequired()
	_, WxUser := request.GetWxUser()
	houses := db_manager.GetUserHouses(WxUser.Id)
	request.RestFulSuccess(map[string]interface{}{"houses": view_manager.GetHouseListInfo(houses)}, "")
}
