package db_manager

import (
	"fmt"
	"rent_backend/models"
)

func GetHouseByQuery(city string, orderBy string, limit int, offset int) (houses []models.HouseModel) {
	var HouseModel models.HouseModel
	qs := models.OrmManager.QueryTable(HouseModel)
	if orderBy != "" {
		qs.OrderBy(orderBy)
	}
	fmt.Println("city", city)
	_, _ = qs.Filter("city__exact", city).Limit(limit, offset).All(&houses)
	return houses
}
