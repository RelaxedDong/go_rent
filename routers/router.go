package routers

import (
	"github.com/astaxie/beego"
	"rent_backend/controllers/account"
	"rent_backend/controllers/house"
	"rent_backend/middleware"
)

func init() {
	middleware.CheckLogin()
	beego.Router("/api/account/login", &account.Controller{}, "Post:Login")
	beego.Router("/api/account/user_info", &account.Controller{}, "Post:BindUserInfo")
	beego.Router("/api/account/edit_info", &account.Controller{}, "Get:UserInfo")
	beego.Router("/api/account/edit_info", &account.Controller{}, "Post:EditUserInfo")
	// house
	beego.Router("/api/house/city_conf", &house.Controller{}, "Get:CityListConf")
	beego.Router("/api/house/index", &house.Controller{}, "Get:HouseIndex")
	beego.Router("/api/house/detail/:house_id([0-9]+)", &house.Controller{}, "Get:HouseDetail")
}
