package routers

import (
	"github.com/astaxie/beego"
	"rent_backend/controllers/account"
)

func init() {
	beego.Router("/login", &account.AccountController{})
}
