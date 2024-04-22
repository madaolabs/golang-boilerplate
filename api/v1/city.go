package v1

import (
	"github.com/gin-gonic/gin"
	"golang-boilerpalte/internal/service"
)

func AddCityRouter(group *gin.RouterGroup) {
	group.GET("/city/list", service.GetCitiesList)
}
