package server

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DBIns *gorm.DB

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(viper.GetString("dbPG.url")), &gorm.Config{})
	if err != nil {
		log.Fatal("err: ", err)
	}
	sqlDb, sqlErr := db.DB()
	if sqlErr != nil {
		log.Fatal("sqlErr", sqlErr)
	}
	//defer sqlDb.Close()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(10)
	fmt.Println("Connect DB Success")
	DBIns = db
	return db
}
