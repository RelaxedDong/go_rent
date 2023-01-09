package account

import (
	"rent_backend/controllers"
	accountform "rent_backend/controllers/account/form"
	UserDbManager "rent_backend/controllers/account/manager/db_manager"
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
	openId, errMsg := weixin.GetUserOpenidAndSessionKey(req.Code)
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
	// todo: is_superuser, finish_user_info字段舍弃？
	data["user_id"] = account.Id
	request.RestFulSuccess(data, "")
}

// BindUserInfo 这里是用户点击授权绑定处理，如果当前用户是第一次授权绑定，则创建一个新用户
func (request *Controller) BindUserInfo() {
	WxUser := request.Ctx.Input.GetData("WxUser")
	// 不存在就创建一个用户
	if WxUser == nil {
		// 断言值是给定的类型
		OpenId := request.Ctx.Input.GetData("openId").(string)
		var req accountform.UserInfoForm
		req.OpenId = OpenId
		request.RequestJsonFormat(&req)
		UserDbManager.GetOrCreateUser(req)
	}
	request.RestFulSuccess(make(map[string]interface{}), "")
}
