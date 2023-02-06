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

type EditUserInfoForm struct {
	Wechat string `json:"wechat" validate:"required|minLen:5|maxLen:30" message:"required: 请输入微信号|min:微信号错误|max:微信号错误"`
	Phone  string `json:"phone" validate:"required|minLen:11|maxLen:11" message:"required: 请输入手机号|min:手机号错误|max:手机号错误"`
	//Age       uint8  `json:"age" validate:"required|min:0|max:111" message:"required: 请输入年龄|min:年龄错误|max:年龄错误"`
	//Signature string `json:"signature" validate:"maxLen:100" message:"maxLen: 签名过长，请重新输入"`
}
type OperationForm struct {
	OperationType string `json:"operation_type" validate:"required" message:"required: 请传递操作类型"`
	HouseId       int64  `json:"houseId"`
}
type BindPhoneForm struct {
	Iv            string `json:"iv" validate:"required"`
	EncryptedData string `json:"encryptedData" validate:"required"`
}
type OperationDeleteForm struct {
	OperationType string `json:"op_type" validate:"required" message:"required: 请传递操作类型"`
	HouseId       int64  `json:"house_id"`
	ClearAll      bool   `json:"clear_all"`
}

// Messages 您可以自定义验证器错误消息
//func (f Login) Messages() map[string]string {
//	return validate.MS{
//		//"required": "这个字段说必须传递的!",
//		"Code.required": "code is required~",
//	}
//}
