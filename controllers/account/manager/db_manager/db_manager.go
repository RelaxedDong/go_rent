package UserDbManager

import (
	"errors"
	accountform "rent_backend/controllers/account/form"
	"rent_backend/models"
	"strconv"
	"time"
)

func GetUserByOpenId(openId string) (user models.AccountModel, err error) {
	if openId == "" {
		return user, errors.New("openId不能为空")
	}
	qs := models.OrmManager.QueryTable(user)
	err = qs.Filter("openid__exact", openId).One(&user)
	if err != nil {
		return user, errors.New("用户不存在")
	}
	return user, nil
}

func GetOrCreateUser(userInfo accountform.UserInfoForm) {
	user := models.AccountModel{
		OpenId:    userInfo.OpenId,
		NickName:  userInfo.NickName,
		AvatarUrl: userInfo.AvatarUrl,
		Province:  userInfo.Province,
		City:      userInfo.City,
		Gender:    strconv.Itoa(int(userInfo.Gender)),
		Phone:     userInfo.Phone,
		LastLogin: time.Now(),
	}
	// 是否是新创建的，创建的id，错误
	_, _, _ = models.OrmManager.ReadOrCreate(&user, "OpenId")
}
