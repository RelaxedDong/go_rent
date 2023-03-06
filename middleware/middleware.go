package middleware

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	UserDbManager "rent_backend/controllers/account/manager/db_manager"
	"rent_backend/utils/jwt"
)

var CookieMaxAge = 60 * 60 * 24 * 14

func CheckLogin() {
	//登录认证中间件过滤器
	var login = func(ctx *context.Context) {
		JwtToken := ctx.Input.GetData("token")
		if JwtToken != "" {
			if Claims, parseTokenErr := jwt.ParseToken(JwtToken.(string)); parseTokenErr == nil {
				ctx.Input.SetData("openId", Claims.OpenId)
				if account, userError := UserDbManager.GetUserByOpenId(Claims.OpenId); userError == nil {
					// 给context对象绑定一个 AccountModel 对象，便于在视图里面直接获取到
					ctx.Input.SetData("WxUser", account)
				}
			}
		}
	}
	// 登录过滤器
	beego.InsertFilter("/*", beego.BeforeRouter, login, false)
}

func ProcessRequest() {
	var SetCookie = func(ctx *context.Context) {
		JwtToken := ctx.Input.Header("token") // 小程序来源
		if JwtToken == "" {
			JwtToken = ctx.GetCookie("token") // 从cookie中拿
			fmt.Println("from cookie", JwtToken)
			if JwtToken == "" {
				JwtToken = ctx.Input.Query("jwt_token") // 小程序webview传递到h5
			}
		}
		ctx.Input.SetData("token", JwtToken)
		if JwtToken != "" {
			ctx.SetCookie("token", JwtToken, CookieMaxAge, "/")
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, SetCookie, false)
}
