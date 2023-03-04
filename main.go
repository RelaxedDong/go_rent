package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/joho/godotenv"
	"path/filepath"
	"rent_backend/models"
	_ "rent_backend/routers"
	_ "rent_backend/template_func"
	"rent_backend/utils"
)

func InitEnv() {
	mode := beego.AppConfig.String("runmode")
	env := filepath.Join("conf", mode+".env")
	configs, _ := godotenv.Read(env)
	for key, value := range configs {
		beego.AppConfig.Set(key, value)
	}
}

func init() {
	InitEnv()
	models.InitOrmConfig()
	logs.Info("logfilepath is" + beego.AppConfig.String("logfilepath"))
	filename := `{"filename": "{filename}","separate":["error", "debug"],"daily":true,"maxdays":7,"color":true}`
	path := utils.FormatString(filename, map[string]interface{}{
		"filename": beego.AppConfig.String("logfilepath"),
	})
	logs.SetLogger(logs.AdapterMultiFile, path)
	logs.EnableFuncCallDepth(true)
	logs.Async()
	if beego.AppConfig.String("runmode") != "dev" {
		//logs.GetBeeLogger().DelLogger("console")              // 删除控制台输出，参数是某个引擎
		logs.GetBeeLogger().SetLevel(logs.LevelInfo) // 日志级别拦截：只会展示当前级别-比当前更高的级别
	}
}

func main() {
	beego.SetViewsPath("templates") // app config里面也可以配置：viewspath=templates, 默认main里面优先
	beego.SetStaticPath("/static", "static")
	// logs.Async(1e3)
	beego.Run()
}
