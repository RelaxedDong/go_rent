package account

type Login struct {
	Code string `json:"code" validate:"required"`
}

// Messages 您可以自定义验证器错误消息
//func (f Login) Messages() map[string]string {
//	return validate.MS{
//		//"required": "这个字段说必须传递的!",
//		"Code.required": "code is required~",
//	}
//}
