package routers

import (
	"github.com/astaxie/beego"
	"rent_backend/controllers/account"
	"rent_backend/middleware"
)

func init() {
	middleware.CheckLogin()
	beego.Router("/api/account/login", &account.Controller{}, "Post:Login")
	beego.Router("/api/account/user_info", &account.Controller{}, "Post:BindUserInfo")
}
