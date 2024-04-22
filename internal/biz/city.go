package biz

import (
	"golang-boilerpalte/internal/data"
	"golang-boilerpalte/internal/server"
)

func GetCitiesList() []data.City {
	//var limitI, _ = strconv.Atoi(limit)
	//var pageI, _ = strconv.Atoi(page)
	var cities []data.City
	server.DBIns.Find(&cities)

	return cities

}

func GetCityDetail(conds *data.City) data.City {
	var cityInfo data.City
	server.DBIns.Find(&cityInfo, &conds)
	return cityInfo
}
