package view_manager

import (
	"encoding/json"
	"reflect"
	"rent_backend/consts"
	"rent_backend/models"
	"rent_backend/utils"
	"rent_backend/utils/datetime"
	"strconv"
	"strings"
)

func GetHouseListInfo(houses []models.HouseModel, excludeIds ...[]int64) []interface{} {
	houseInfoList := []interface{}{}
	exclude := []int64{}
	if len(excludeIds) > 0 {
		exclude = excludeIds[0]
	}
	for _, house := range houses {
		isInExclude := false
		if len(exclude) > 0 {
			for _, i3 := range exclude {
				if house.Id == i3 {
					isInExclude = true
					break
				}
			}
		}
		if isInExclude {
			continue
		}
		houseInfoList = append(houseInfoList, BuildHouseInfo(house))
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
	// 这里重新赋值返回，便于后续过滤操作
	houseTypeValue := house.HouseType
	ApartmentValue := house.Apartment
	house.Apartment = consts.ApartMentTypeMap[house.Apartment]
	house.HouseType = consts.RentTypeMap[house.HouseType]
	houseInfo := map[string]interface{}{
		"id":               house.Id,
		"title":            house.Title,
		"desc":             house.Desc,
		"address":          house.Address,
		"region":           house.Region,
		"province":         house.Province,
		"city":             house.City,
		"status":           house.Status,
		"status_verbose":   statusVerbose,
		"house_type":       house.HouseType,
		"house_type_value": houseTypeValue,
		"apartment_value":  ApartmentValue,
		"apartment":        house.Apartment,
		"short_rent":       ShortRent,
		"can_short_rent":   house.CanShortRent,
		"video_url":        house.VideoUrl,
		"user_id":          house.Publisher.Id,
		"longitude":        house.Longitude,
		"latitude":         house.Latitude,
		"imgs":             Imgs,
		"storey":           house.Storey,
		"area":             house.Area,
		"price":            house.Price,
		"facilities":       Facilities,
		"create_time":      datetime.TimeToStr(house.CreatedTime),
		"last_modified":    datetime.TimeToStr(house.LastModified),
		"is_delicate":      house.IsDelicate,
		"tags":             BuildHouseTags(house),
		"tags_conf":        BuildHouseElement(house),
	}
	publisher := house.Publisher
	houseInfo["publisher"] = map[string]string{
		"nickname":   publisher.NickName,
		"avatarUrl":  publisher.AvatarUrl,
		"gender":     publisher.Gender,
		"last_login": datetime.GetLastLoginActiveText(publisher.LastLogin),
		"wechat":     publisher.Wechat,
		//"phone":      publisher.Phone,
	}
	facilitiesConfList := []map[string]string{}
	for _, i2 := range houseInfo["facilities"].([]string) {
		data := consts.FacilityMap[i2]
		facilitiesConfList = append(facilitiesConfList, data)
	}
	houseInfo["facilities_conf"] = facilitiesConfList
	return houseInfo
}

var TagElementList = map[string]string{
	"Price":        "房租",
	"Region":       "区域",
	"Subway":       "地铁",
	"OwnerRent":    "直租",
	"Direction":    "朝向",
	"Apartment":    "房型",
	"Area":         "面积",
	"Storey":       "楼层",
	"BedroomType":  "卧室",
	"CanShortRent": "可短租",
	"PayMethod":    "支付方式",
}

func BuildHouseElement(house models.HouseModel) map[string]interface{} {
	result := map[string]interface{}{}
	tagsArray := []string{
		"Price",
		"Storey",
		"Apartment",
		//"Direction",
		//"OwnerRent",
		//"BedroomType",
		//"PayMethod",
		"Region",
		"Area",
		"Subway",
		"CanShortRent",
	}
	v := reflect.ValueOf(house)
	for _, key := range tagsArray {
		value := v.FieldByName(key).Interface()
		switch o := value.(type) {
		case bool:
			if o {
				value = "是"
			} else {
				value = "否"
			}
		}
		result[utils.ToSnakeCase(key)] = map[string]interface{}{
			"name":  TagElementList[key],
			"value": value,
		}
	}
	return result
}

func BuildHouseTags(house models.HouseModel) []string {
	tags := []string{house.HouseType, house.Apartment}
	if house.CanShortRent {
		tags = append(tags, "可短租")
	}
	// todo: 判断视频链接
	if house.VideoUrl != "" {
		tags = append(tags, "视频实拍")
	}
	return tags
}

func GetHomePageConfig(banners []models.BannerModel) ([]interface{}, []interface{}) {
	bannerList := []interface{}{}
	iconList := []interface{}{}
	for _, banner := range banners {
		dataInfo := BuildBannerDict(banner)
		switch banner.Position {
		case "banner":
			bannerList = append(bannerList, dataInfo)
		case "icon":
			iconList = append(iconList, dataInfo)
		}
	}
	return bannerList, iconList
}

func BuildBannerDict(banner models.BannerModel) interface{} {
	return map[string]interface{}{
		"url":      banner.Url,
		"desc":     banner.Desc,
		"navigate": banner.Navigate,
		"is_show":  banner.IsShow,
	}
}

func BuildSearchFilters(apartmentList interface{}, houseTypeList interface{}) interface{} {
	facilityComponents := []map[string]string{}
	for key, value := range consts.FacilityMap {
		facilityComponents = append(facilityComponents, map[string]string{
			"name": value["name"], "value": key,
		})
	}
	return map[string]interface{}{
		"dropDownMenuTitle": []string{"综合排序", "位置", "价格", "筛选"},
		"dropDownMenuFirstData": []map[string]string{
			{"name": "综合排序", "value": "default"},
			{"name": "低价排序", "value": "price"},
			{"name": "最新发布", "value": "time"},
			//{"name": "附近房源", "value": "distance"},
		},
		"dropDownMenuThirdData": []map[string]string{
			{"name": "1500元以下", "value": "0-1500"},
			{"name": "1500-2000元", "value": "1500-2000"},
			{"name": "2000-2500元", "value": "2000-2500"},
			{"name": "2500-3000元", "value": "2500-3000"},
			{"name": "3000-3500元", "value": "3000-3500"},
			{"name": "3500-4000元", "value": "3500-4000"},
		},
		"dropDownMenuFourthData": []map[string]interface{}{
			{"name": "房型", "value": "apartment", "components": apartmentList},
			{"name": "整租/合租", "value": "house_type", "components": houseTypeList},
			{"name": "设施", "value": "facility", "components": facilityComponents},
		},
	}
}

var SortConf = map[string]string{
	"default":  "-view_count",
	"price":    "price",
	"time":     "-created_time",
	"distance": "distance",
}

func GetHouseFilterBySearch(filter map[string]interface{}) map[string]interface{} {
	var region, subways, sortBy string
	var startPrice, endPrice uint64
	if searchSort, ok := filter["1"]; ok {
		sortBy, _ = SortConf[searchSort.(string)]
	}
	if regionSubwayConf, ok := filter["2"]; ok {
		regionSubwayTuple := regionSubwayConf.([]interface{})
		region = regionSubwayTuple[0].(string)
		subwaysList := regionSubwayTuple[1].([]interface{})
		// 地铁站搜索，目前只支持一个
		if len(subwaysList) > 0 {
			subways = subwaysList[0].(string)
		}
	}
	if priceConf, ok := filter["3"]; ok {
		// 如果是空字符串，priceRange=[""]
		priceRange := strings.Split(priceConf.(string), "-")
		if len(priceRange) > 1 {
			startPrice, _ = strconv.ParseUint(priceRange[0], 10, 64)
			endPrice, _ = strconv.ParseUint(priceRange[1], 10, 64)
		}
	}
	var apartments []interface{}
	var houseTypes []interface{}
	var facilitiesList []interface{}
	if sourceConf, ok := filter["4"]; ok {
		if apartmentsFilter, withApartment := sourceConf.(map[string]interface{})["0"]; withApartment {
			apartments = apartmentsFilter.([]interface{})
		}
		if HouseTypeFilter, withHouseType := sourceConf.(map[string]interface{})["1"]; withHouseType {
			houseTypes = HouseTypeFilter.([]interface{})
		}
		if facilitiesFilter, withHouseType := sourceConf.(map[string]interface{})["2"]; withHouseType {
			facilitiesList = facilitiesFilter.([]interface{})
		}
	}
	result := map[string]interface{}{
		"sort_by":        sortBy,
		"region":         region,
		"subways":        subways,
		"startPrice":     startPrice,
		"endPrice":       endPrice,
		"apartments":     utils.InterFaceListToStringList(apartments),
		"facilitiesList": utils.InterFaceListToStringList(facilitiesList),
		"houseTypes":     utils.InterFaceListToStringList(houseTypes),
	}
	return result
}
