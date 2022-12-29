package account

import (
	"encoding/json"
	"rent_backend/controllers"
	_ "rent_backend/models"
)

type AccountController struct {
	controllers.BaseController
}

func (request *AccountController) Get() {

}

func (request *AccountController) Post() {
	RequestBody := request.Ctx.Input.RequestBody
	var body Login
	err := json.Unmarshal(RequestBody, &body)
	if err != nil {
		request.RestFulParamsError(map[string]interface{}{}, "json解析错误: "+err.Error())
	}
	data := map[string]interface{}{
		"token":            "",
		"user_id":          1,
		"is_superuser":     "",
		"finish_user_info": "",
	}
	request.RestFulSuccess(data, "")
}
