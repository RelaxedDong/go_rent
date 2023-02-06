package db_manager

import (
	"fmt"
	"rent_backend/models"
	"rent_backend/utils"
	"strings"
)

func GetNearByHouses(city string, lng float64, lat float64, miles uint8) (houses []models.HouseModel) {
	sql := "SELECT *, ((ACOS(SIN({lat} * PI() / 180) * SIN(latitude * PI() / 180) + COS({lat} * PI() / 180) * COS(latitude * PI() / 180) * COS(({lng} - longitude) * PI() / 180)) * 180 / PI()) * 60 * 1.1515) AS distance FROM house having city='{city}' and distance<={miles} ORDER BY distance ASC"
	query := utils.FormatString(sql, map[string]interface{}{"lng": lng, "lat": lat, "miles": miles, "city": city})
	models.OrmManager.Raw(query).QueryRows(&houses)
	return houses
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
	qs := models.OrmManager.QueryTable(HouseModel).Filter("city__exact", city)
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

	_, err := qs.RelatedSel("publisher").Limit(limit, offset).All(&houses)
	fmt.Println("err", err)
	return houses
}

func GetHouseById(id int64) (house models.HouseModel, err error) {
	qs := models.OrmManager.QueryTable(house)
	err = qs.Filter("id__exact", id).RelatedSel("publisher").One(&house)
	return house, err
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
