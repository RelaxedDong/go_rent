package main

import (
	"github.com/astaxie/beego"
	_ "rent_backend/routers"
)

func main() {
	beego.SetViewsPath("templates") // app config里面也可以配置：viewspath=templates, 默认main里面优先
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
