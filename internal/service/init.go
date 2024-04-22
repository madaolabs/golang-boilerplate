package service

import (
	"golang-boilerpalte/internal/data"
	"golang-boilerpalte/internal/server"
)

func InitDatabase() {
	server.DBIns.AutoMigrate(&data.City{})
}
