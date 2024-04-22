package main

import (
	"github.com/spf13/viper"
	"golang-boilerpalte/api"
	"golang-boilerpalte/configs"
	"golang-boilerpalte/internal/server"
)

func main() {
	configs.InitConfig()
	//server.ConnectDB()
	//service.InitDatabase()
	r := server.HttpClient()

	api.AddV1Router(r)
	r.Run(":" + viper.GetString("server.portHttp"))
}
