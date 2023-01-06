package account

import (
	"fmt"
	"rent_backend/controllers"
	UserDbManager "rent_backend/controllers/account/manager/db_manager"
	_ "rent_backend/models"
	"rent_backend/third_party_service/weixin"
	"rent_backend/utils/jwt"
)

type AccountController struct {
	controllers.BaseController
}

func (request *AccountController) Get() {

}

func (request *AccountController) Post() {
	var req Login
	var data map[string]interface{}
	request.RequestJsonFormat(&req)
	openId, errMsg := weixin.GetUserOpenidAndSessionKey(req.Code)
	fmt.Println("openId", openId)
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