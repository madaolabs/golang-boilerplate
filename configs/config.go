package configs

const ConfigPath = "/configs"
const ENV = "dev"

func InitConfig() {
	InitViper(ENV, "yml", ConfigPath)
}
