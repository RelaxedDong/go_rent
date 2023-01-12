package view_manager

import (
	"encoding/json"
	"rent_backend/consts"
	"rent_backend/models"
	"rent_backend/utils/datetime"
)

func GetHouseListInfo(houses []models.HouseModel) (houseInfoList []interface{}) {
	for i := 0; i < len(houses); i++ {
		houseInfoList = append(houseInfoList, BuildHouseInfo(houses[i]))
	}
	return houseInfoList
}

func BuildHouseInfo(house models.HouseModel) interface{} {
	var statusVerbose = consts.HouseStatusMap[house.Status]
	var ShortRent string
	// var Imgs []string，Imgs 会返回nil
	Imgs := []string{}
	Facilities := []string{}
	if house.Status == consts.CHECKFAIL {
		statusVerbose = house.FailReason
	}
	if house.CanShortRent {
		ShortRent = "可短租"
	}
	json.Unmarshal([]byte(house.Imgs), &Imgs)
	json.Unmarshal([]byte(house.Facilities), &Facilities)
	return map[string]interface{}{
		"id":             house.Id,
		"title":          house.Title,
		"address":        house.Address,
		"region":         house.Region,
		"city":           house.City,
		"status":         house.Status,
		"status_verbose": statusVerbose,
		"house_type":     consts.RentTypeMap[house.HouseType],
		"apartment":      consts.ApartMentTypeMap[house.Apartment],
		"short_rent":     ShortRent,
		"video_url":      house.VideoUrl,
		"user_id":        house.Publisher.Id,
		"longitude":      house.Longitude,
		"latitude":       house.Latitude,
		"imgs":           Imgs,
		"facilities":     Facilities,
		"create_time":    datetime.DateTimeToStr(house.CreatedTime),
		"last_modified":  datetime.DateTimeToStr(house.UpdatedTime),
		"is_delicate":    house.IsDelicate,
	}
}
