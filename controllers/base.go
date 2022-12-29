package controllers

import (
	"github.com/astaxie/beego"
	"rent_backend/consts"
)

type BaseController struct {
	beego.Controller
}

func (response *BaseController) WriteResponse(data interface{}, msg string, default_code int, customCode ...int) {
	var code = default_code
	if len(customCode) > 0 {
		code = customCode[0]
	}
	response.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	response.ServeJSON()
	response.StopRun()
}

func (response *BaseController) RestFulSuccess(data interface{}, msg string, args ...int) {
	// status_code 使用可变参数语法糖
	response.WriteResponse(data, msg, consts.STATUS_CODE_200, args...)
}

func (response *BaseController) RestFulParamsError(data interface{}, msg string, args ...int) {
	response.WriteResponse(data, msg, consts.STATUS_CODE_400, args...)
}

//type UserController struct {
//	beego.Controller
//}
//
//type UserStruct struct {
//	Id   int
//	Name string
//	Age  int
//}
//
//func (c *MainController) Get() {
//	// name := c.GetString("name")
//	elasticsearch_addr := beego.AppConfig.String("123")
//	if elasticsearch_addr != "" {
//		fmt.Println("elasticsearch_addr is ", elasticsearch_addr)
//	} else {
//		fmt.Println("elasticsearch_addr None")
//	}
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}
//
//func (c *UserController) Get() {
//	// ?xxx=xxx
//	//id := c.GetString("id") // POST请求值也能拿到。。
//	//user := c.Input().Get("name") // POST请求值也能拿到。。
//
//	//age, _ = c.GetInt64("age")
//	//c.GetBool("is_true")
//	//c.GetFloat("money")
//
//	// /xxxxxx
//	//id := c.GetString(":id")
//	//c.Input().Get("name") // 获取不到
//	id := c.Ctx.Input.Param(":id")
//	fmt.Println("from front user", id)
//
//	// 使用指针
//	c.Data["name"] = "donghao"
//	c.TplName = "user.html"
//	//arrs := []int{1, 2, 3, 4, 5}
//	user1 := UserStruct{1, "donghao", 24}
//	user2 := UserStruct{2, "tanyajuan", 24}
//	users := []UserStruct{user1, user2}
//	//c.Data["users"] = users
//	c.Data["users"] = users
//
//	// key value对
//	//var mapa map[string]string
//	//var mapa = make(map[string]string)
//	//mapa := map[string]string{"Name": "donghao"}
//	//fmt.Println("mapa", mapa["Name"])
//	//age, ok := mapa["Age"]
//	//fmt.Printf("%T\n", age, ok)
//
//}
