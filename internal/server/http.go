package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HttpClient() *gin.Engine {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	r.Use(cors.New(corsConfig))

	return r
}
