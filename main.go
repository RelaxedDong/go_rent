package main

import (
	"github.com/astaxie/beego"
	_ "rent_backend/init"
	_ "rent_backend/routers"
	_ "rent_backend/template_func"
)

func main() {
	beego.SetViewsPath("templates") // app config里面也可以配置：viewspath=templates, 默认main里面优先
	beego.SetStaticPath("/static", "static")
	// logs.Async(1e3)
	beego.Run()
}
