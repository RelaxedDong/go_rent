package UserDbManager

import (
	"errors"
	"rent_backend/models"
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
