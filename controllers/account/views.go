package account

import (
	"fmt"
	"github.com/gookit/validate"
	"rent_backend/controllers"
	_ "rent_backend/models"
)

type AccountController struct {
	controllers.BaseController
}

func (request *AccountController) Get() {

}

func (request *AccountController) Post() {
	var req Login
	request.RequestJsonFormat(&req)
	// 参数校验
	v := validate.Struct(req)
	if !v.Validate() {
		request.RestFulParamsError("参数校验: " + v.Errors.One())
	}
	fmt.Println("body", req)
	data := map[string]interface{}{
		"token":            "",
		"user_id":          1,
		"is_superuser":     "",
		"finish_user_info": "",
	}
	request.RestFulSuccess(data, "")
}
