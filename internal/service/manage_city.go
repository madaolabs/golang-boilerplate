package service

import (
	"github.com/gin-gonic/gin"
	"golang-boilerpalte/internal/biz"
	"net/http"
)

type GetCityParams struct {
	PageSize int `form:"pageSize" binding:"required"`
	Page     int `form:"page" binding:"required"`
}

func GetCitiesList(c *gin.Context) {
	cities := biz.GetCitiesList()

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data": gin.H{
			"addresses": cities,
		},
	})
}
