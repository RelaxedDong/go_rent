package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gookit/validate"
	"rent_backend/consts"
)

type BaseController struct {
	beego.Controller
}

func (req *BaseController) RequestJsonFormat(structBody interface{}) {
	// 这里只会判断类型，因为需要把json数据解析到结构体里面
	err := json.Unmarshal(req.Ctx.Input.RequestBody, structBody)
	if e, ok := err.(*json.UnmarshalTypeError); ok {
		req.RestFulParamsError(e.Field + "(" + e.Type.String() + ")" + " 字段类型错误 ->" + e.Value)
	}
	// 使用gookit/validate做校验
	v := validate.Struct(structBody)
	if !v.Validate() {
		req.RestFulParamsError(v.Errors.One())
	}
}

func (req *BaseController) WriteResponse(data interface{}, msg string, defaultCode int, customCode ...int) {
	var code = defaultCode
	if len(customCode) > 0 {
		code = customCode[0]
	}
	req.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	req.ServeJSON()
	req.StopRun()
}

func (req *BaseController) RestFulSuccess(data interface{}, msg string, args ...int) {
	// status_code 使用可变参数语法糖
	req.WriteResponse(data, msg, consts.STATUS_CODE_200, args...)
}

func (req *BaseController) RestFulParamsError(msg string, args ...int) {
	req.WriteResponse(map[string]interface{}{}, msg, consts.STATUS_CODE_400, args...)
}
