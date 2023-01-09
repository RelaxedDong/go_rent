package accountform

type LoginForm struct {
	Code string `json:"code" validate:"required"`
}

type UserInfoForm struct {
	NickName string `json:"nickName"`
	// 0: 未知 1：男 2：女
	Gender    uint8  `json:"gender" validate:"min:0|max:2"`
	AvatarUrl string `json:"avatarUrl"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Phone     string `json:"phone"`
	OpenId    string `json:"OpenId"`
}

// Messages 您可以自定义验证器错误消息
//func (f Login) Messages() map[string]string {
//	return validate.MS{
//		//"required": "这个字段说必须传递的!",
//		"Code.required": "code is required~",
//	}
//}
