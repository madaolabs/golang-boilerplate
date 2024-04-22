package configs

import (
	"github.com/spf13/viper"
	"os"
)

func InitViper(name string, viperType string, path string) {
	workDir, _ := os.Getwd()
	viper.SetConfigName(name)
	viper.SetConfigType(viperType)
	viper.AddConfigPath(workDir + path)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
