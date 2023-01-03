package account

type Login struct {
	Code int    `json:"code" validate:"required|minLen:3" message:"minLen: code 长度不能小于3"`
	Name string `json:"name" validate:"required|minLen:3" message:"minLen: name 长度不能小于3"`
}

// Messages 您可以自定义验证器错误消息
//func (f Login) Messages() map[string]string {
//	return validate.MS{
//		//"required": "这个字段说必须传递的!",
//		"Code.required": "code is required~",
//	}
//}
