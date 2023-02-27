package UserDbManager

import (
	"errors"
	"rent_backend/consts"
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
		return user, errors.New(consts.ErrorMsgAccountNotExists)
	}
	return user, nil
}

func GetOrCreateUser(userInfo accountform.UserInfoForm) (IsNew bool, UserId int64) {
	user := models.AccountModel{
		OpenId:       userInfo.OpenId,
		NickName:     userInfo.NickName,
		AvatarUrl:    userInfo.AvatarUrl,
		Province:     userInfo.Province,
		City:         userInfo.City,
		Gender:       strconv.Itoa(int(userInfo.Gender)),
		Phone:        userInfo.Phone,
		LastLogin:    time.Now(),
		FromPlatform: "weixin",
	}
	// 是否是新创建的，创建的id，错误
	IsNew, UserId, _ = models.OrmManager.ReadOrCreate(&user, "OpenId")
	return IsNew, UserId
}

func UpdateUserInfo(userInfo models.AccountModel, updatedFields []string, updateLoginTime bool) {
	if updateLoginTime {
		userInfo.LastLogin = time.Now()
		updatedFields = append(updatedFields, "last_login")
	}
	models.OrmManager.Update(&userInfo, updatedFields...)
}

func GetOrCreateUserFavor(house models.HouseModel, account models.AccountModel) (isNew bool, RecordId int64) {
	Collect := models.CollectModel{
		Publisher: &account,
		House:     &house,
	}
	// 是否是新创建的，创建的id，错误
	isNew, RecordId, _ = models.OrmManager.ReadOrCreate(&Collect, "Publisher", "House")
	return isNew, RecordId
}

func GetOrCreateUserHistory(house models.HouseModel, account models.AccountModel) (isNew bool, RecordId int64) {
	history := models.ViewHistory{
		Publisher: &account,
		House:     &house,
	}
	isNew, RecordId, _ = models.OrmManager.ReadOrCreate(&history, "Publisher", "House")
	return isNew, RecordId
}

func DeleteCollectRecordDb(RecordId int64) {
	models.OrmManager.Delete(&models.CollectModel{Id: RecordId})
}

func DeleteUserCollectRecord(HouseId int64, PublisherId int64) {
	_, _ = models.OrmManager.QueryTable(models.CollectModel{}).Filter("House__id", HouseId).Filter("Publisher__id", PublisherId).Delete()
}

func DeleteHistoryRecordDb(HouseId int64, PublisherId int64) {
	_, _ = models.OrmManager.QueryTable(models.ViewHistory{}).Filter("House__id", HouseId).Filter("Publisher__id", PublisherId).Delete()

}

func DeleteAllHistoryByUserId(Publisher models.AccountModel) {
	models.OrmManager.QueryTable(models.ViewHistory{}).Filter("Publisher", Publisher).Delete()
}

func DeleteAllCollectsByUserId(Publisher models.AccountModel) {
	models.OrmManager.QueryTable(models.CollectModel{}).Filter("Publisher", Publisher).Delete()
}

func GetHousesByIds(Ids []int64) []models.HouseModel {
	houses := []models.HouseModel{}
	if len(Ids) == 0 {
		return houses
	}
	models.OrmManager.QueryTable(models.HouseModel{}).Filter("Id__in", Ids).RelatedSel("publisher").All(&houses)
	return houses
}

func GetUserCollects(UserId int64, offset uint64, limit uint64) []models.HouseModel {
	var collects []models.CollectModel
	models.OrmManager.QueryTable(models.CollectModel{}).Filter("publisher",
		UserId).OrderBy("-created_time").Limit(limit, offset).All(&collects, "House")
	houseIds := []int64{}
	for _, i2 := range collects {
		houseIds = append(houseIds, i2.House.Id)
	}
	return GetHousesByIds(houseIds)
}

func GetUserHistoryList(UserId int64, offset uint64, limit uint64) []models.HouseModel {
	var collects []models.ViewHistory
	models.OrmManager.QueryTable(models.ViewHistory{}).Filter("publisher",
		UserId).OrderBy("-created_time").Limit(limit, offset).All(&collects, "House")
	houseIds := []int64{}
	for _, i2 := range collects {
		houseIds = append(houseIds, i2.House.Id)
	}
	return GetHousesByIds(houseIds)
}

func IsUserCollectHouse(HouseId int64, UserId int64) (isExist bool) {
	qs := models.OrmManager.QueryTable(models.CollectModel{})
	isExist = qs.Filter("house__exact", HouseId).Filter("publisher__exact", UserId).Exist()
	return isExist
}
