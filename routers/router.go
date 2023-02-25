package routers

import (
	"github.com/astaxie/beego"
	"rent_backend/controllers/account"
	"rent_backend/controllers/common"
	"rent_backend/controllers/house/view/house_api"
	"rent_backend/controllers/web"
	"rent_backend/middleware"
)

func init() {
	middleware.CheckLogin()

	beego.Router("/api/account/login", &account.Controller{}, "Post:Login")
	beego.Router("/api/account/user_info", &account.Controller{}, "Post:BindUserInfo")
	beego.Router("/api/account/bind_phone", &account.Controller{}, "Post:BindPhone")
	beego.Router("/api/account/edit_info", &account.Controller{}, "Get:UserInfo")
	beego.Router("/api/account/edit_info", &account.Controller{}, "Post:EditUserInfo")
	beego.Router("/api/account/operation", &account.Controller{}, "Post:Operation")
	beego.Router("/api/account/get_collects", &account.Controller{}, "Get:Collects")
	beego.Router("/api/account/collects_delete", &account.Controller{}, "Post:OperationDelete")
	beego.Router("/api/account/get_publish", &account.Controller{}, "Get:UserPublish")
	// house api
	beego.Router("/api/house/city_conf", &house_api.Controller{}, "Get:CityListConf")
	beego.Router("/api/house/index", &house_api.Controller{}, "Get:HouseIndex")
	beego.Router("/api/house/selects", &common.Controller{}, "Get:Selects")
	beego.Router("/api/house/oss_sign", &common.Controller{}, "Get:GetOssSign")
	beego.Router("/api/house/house_add", &house_api.Controller{}, "Post:HouseAdd")
	beego.Router("/api/house/house_add_check", &house_api.Controller{}, "Get:HouseAddCheck")
	beego.Router("/api/house/house_delete", &house_api.Controller{}, "Post:HouseDelete")
	beego.Router("/api/house/banners", &house_api.Controller{}, "Get:BannerList")
	beego.Router("/api/house/search", &house_api.Controller{}, "Get:SearchHouse")
	beego.Router("/api/house/nearby_houses", &house_api.Controller{}, "Get:NearbyHouses")
	beego.Router("/api/house/detail/:house_id([0-9]+)", &house_api.Controller{}, "Get:HouseDetail")
	// pc/m
	beego.Router("/detail/:house_id([0-9]+)", &web.Controller{}, "Get:HouseDetail")
	beego.Router("/ask_rent", &web.Controller{}, "Get:AskRentIndex")
}
