package middleware

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"rent_backend/consts"
	UserDbManager "rent_backend/controllers/account/manager/db_manager"
	"rent_backend/utils"
	"rent_backend/utils/jwt"
)

var IgnoreLoginRequired = []string{
	"/api/account/login",
	"/api/account/user_info",
	"/api/house/index",
	"/api/house/selects",
	"/api/house/banners",
}

func CheckLogin() {
	//登录认证中间件过滤器
	var login = func(ctx *context.Context) {
		if JwtToken := ctx.Input.Header("token"); JwtToken != "" {
			if Claims, parseTokenErr := jwt.ParseToken(JwtToken); parseTokenErr == nil {
				ctx.Input.SetData("openId", Claims.OpenId)
				if account, userError := UserDbManager.GetUserByOpenId(Claims.OpenId); userError == nil {
					// 给context对象绑定一个 AccountModel 对象，便于在视图里面直接获取到
					ctx.Input.SetData("WxUser", account)
				}
			}
		}
	}

	var loginRequired = func(ctx *context.Context) {
		// 必须要登陆的情况 ---> 需要获取到WxUser对象
		path := ctx.Request.URL.Path
		if !utils.InStringArray(path, IgnoreLoginRequired) {
			account := ctx.Input.GetData("WxUser")
			if account == nil {
				resp, _ := json.Marshal(map[string]interface{}{"code": consts.STATUS_CODE_400,
					"msg": "获取用户失败，请重新授权绑定～",
				})
				ctx.Output.Body(resp)
			}
		}
	}
	// 登录过滤器
	beego.InsertFilter("/*", beego.BeforeRouter, login)
	beego.InsertFilter("/*", beego.BeforeRouter, loginRequired)
}
