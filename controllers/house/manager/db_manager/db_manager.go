package db_manager

import (
	"encoding/json"
	"rent_backend/consts"
	"rent_backend/consts/model_consts"
	houseform "rent_backend/controllers/house/form"
	"rent_backend/models"
	"rent_backend/utils"
	"strings"
)

func GetNearByHouses(city string, lng float64, lat float64, miles uint8) (houses []models.HouseModel) {
	sql := "SELECT *, ((ACOS(SIN({lat} * PI() / 180) * SIN(latitude * PI() / 180) + COS({lat} * PI() / 180) * COS(latitude * PI() / 180) * COS(({lng} - longitude) * PI() / 180)) * 180 / PI()) * 60 * 1.1515) AS distance FROM house having city='{city}' and status='{status}' and distance<={miles} ORDER BY distance ASC"
	query := utils.FormatString(sql, map[string]interface{}{"lng": lng, "lat": lat, "miles": miles, "city": city, "status": model_consts.HOUSE_RENTING})
	models.OrmManager.Raw(query).QueryRows(&houses)
	GetHouseAccounts(&houses)
	return houses
}

func GetHouseAccounts(houses *[]models.HouseModel) []models.HouseModel {
	userIds := []int64{}
	for _, house := range *houses {
		if !utils.IsInInt64Arrray(userIds, house.Publisher.Id) {
			userIds = append(userIds, house.Publisher.Id)
		}
	}
	UserIdToUser := map[int64]models.AccountModel{}
	users := []models.AccountModel{}
	if len(userIds) > 0 {
		models.OrmManager.QueryTable(models.AccountModel{}).Filter("id__in", userIds).All(&users)
		for _, user := range users {
			UserIdToUser[user.Id] = user
		}
	}
	for _, house := range *houses {
		*house.Publisher = UserIdToUser[house.Publisher.Id]
	}
	return *houses
}

func GetHouseByQuery(city string, title string,
	region string,
	houseType []string,
	ApartmentList []string,
	facilitiesList []string,
	priceRange []uint64,
	orderBy string,
	offset uint64, limit int) (houses []models.HouseModel) {
	var HouseModel models.HouseModel
	qs := models.OrmManager.QueryTable(HouseModel).Filter("city__exact", city).Filter("status__exact", model_consts.HOUSE_RENTING)
	if title != "" {
		qs = qs.Filter("title__contains", title)
	}
	if region != "" {
		qs = qs.Filter("region", region)
	}
	if len(houseType) > 0 {
		qs = qs.Filter("house_type__in", houseType)
	}
	if len(ApartmentList) > 0 {
		qs = qs.Filter("apartment__in", ApartmentList)
	}
	if len(priceRange) > 0 && priceRange[1] > 0 {
		qs = qs.Filter("price__gte", priceRange[0]).Filter("price__lte", priceRange[1])
	}
	if len(facilitiesList) > 0 {
		facilitiesFilter := strings.Join(facilitiesList, "|")
		qs = qs.FilterRaw("facilities", "REGEXP '"+facilitiesFilter+"'")
	}
	if orderBy != "" {
		qs = qs.OrderBy(orderBy)
	}

	_, _ = qs.RelatedSel("publisher").Limit(limit, offset).All(&houses)
	return houses
}

func GetHouseById(id int64) (house models.HouseModel, err error) {
	qs := models.OrmManager.QueryTable(house)
	err = qs.Filter("id__exact", id).RelatedSel("publisher").One(&house)
	return house, err
}

func DeleteHouse(house models.HouseModel) (err error) {
	house.Status = consts.DELETED
	_, err = models.OrmManager.Update(&house, "status")
	return err
}

func GetHouseByIdNoPublisher(id int64) (house models.HouseModel, err error) {
	qs := models.OrmManager.QueryTable(house)
	err = qs.Filter("id__exact", id).One(&house)
	return house, err
}

func GetUserHouses(UserId int64) (houses []models.HouseModel) {
	models.OrmManager.QueryTable(models.HouseModel{}).Filter("Publisher__Id", UserId).RelatedSel("publisher").All(&houses)
	return houses
}

func GetUserStatistics(UserId int64) interface{} {
	var allViewCnt, housesCnt uint32
	houses := []models.HouseModel{}
	houseTitles, houseViewCnt := []string{}, []uint32{}
	models.OrmManager.QueryTable(models.HouseModel{}).Filter("Publisher__Id", UserId).All(&houses)
	for _, house := range houses {
		allViewCnt += house.ViewCount
		housesCnt += 1
		houseTitles = append(houseTitles, house.Title)
		houseViewCnt = append(houseViewCnt, house.ViewCount)
	}
	return map[string]interface{}{
		"allViewCnt":   allViewCnt,
		"housesCnt":    housesCnt,
		"houseTitles":  houseTitles,
		"houseViewCnt": houseViewCnt,
	}
}

func getHouseByForm(form houseform.HouseAddForm) models.HouseModel {
	jsonImages, _ := json.Marshal(form.Images)
	FacilityList, _ := json.Marshal(form.FacilityList)
	ProvinceCityRegion := form.ProvinceCityRegion
	province, city, region := ProvinceCityRegion[0], ProvinceCityRegion[1], ProvinceCityRegion[2]
	// todo: 内聚到model层
	var houseType = "1" // 整租
	if form.HouseType == "合租" {
		houseType = "2"
	}
	house := models.HouseModel{
		Title:        form.Title,
		Price:        form.Price,
		Storey:       form.Storey,
		Area:         form.Area,
		Desc:         form.Desc,
		HouseType:    houseType,
		Apartment:    form.Apartment,
		Address:      form.Address,
		Latitude:     form.Latitude,
		Longitude:    form.Longitude,
		Province:     province,
		City:         city,
		Region:       region,
		Imgs:         string(jsonImages),
		Facilities:   string(FacilityList),
		CanShortRent: form.ShortRent,
		Status:       model_consts.HOUSE_CHECKING,
	}
	return house
}

func CreateHouse(form houseform.HouseAddForm, Publisher models.AccountModel) (err error) {
	house := getHouseByForm(form)
	house.Publisher = &Publisher
	_, err = models.OrmManager.Insert(&house)
	return err
}

func UpdateHouse(HouseId int64, form houseform.HouseAddForm, Publisher models.AccountModel) (err error) {
	house := getHouseByForm(form)
	house.Id = HouseId
	house.Publisher = &Publisher
	_, err = models.OrmManager.Update(&house)
	return err
}

func IncreaseHouseViewCnt(house *models.HouseModel) {
	house.ViewCount += 1
	models.OrmManager.Update(house)
}

func GetBannerByQuery(city string, positions []string) (banners []models.BannerModel, err error) {
	var banner models.BannerModel
	qs := models.OrmManager.QueryTable(banner).Filter("position__in", positions)
	if city != "" {
		qs.Filter("city__exact", city)
	}
	_, err = qs.OrderBy("-priority").All(&banners)
	return banners, err
}
