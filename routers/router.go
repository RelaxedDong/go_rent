package routers

import (
	"github.com/astaxie/beego"
	"rent_backend/controllers/account"
)

func init() {
	beego.Router("/api/account/login", &account.AccountController{})
}
