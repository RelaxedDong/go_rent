package init

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/joho/godotenv"
	"path/filepath"
	"rent_backend/models"
	"rent_backend/utils"
)

func initEnv() {
	mode := beego.AppConfig.String("runmode")
	env := filepath.Join("./conf", mode+".env")
	configs, _ := godotenv.Read(env)
	for key, value := range configs {
		beego.AppConfig.Set(key, value)
	}
}

func initLogger() {
	logs.Info("logfilepath is" + beego.AppConfig.String("logfilepath"))
	filename := `{"filename": "{filename}","separate":["error", "debug"],"daily":true,"maxdays":7,"color":true}`
	path := utils.FormatString(filename, map[string]interface{}{
		"filename": beego.AppConfig.String("logfilepath"),
	})
	logs.SetLogger(logs.AdapterMultiFile, path)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}

func init() {
	initEnv()
	initLogger()
	models.InitOrmConfig()

	if beego.AppConfig.String("runmode") != "dev" {
		//logs.GetBeeLogger().DelLogger("console")              // 删除控制台输出，参数是某个引擎
		logs.GetBeeLogger().SetLevel(logs.LevelInfo) // 日志级别拦截：只会展示当前级别-比当前更高的级别
	}
}
