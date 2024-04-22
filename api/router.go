package api

import (
	"github.com/gin-gonic/gin"
	v1 "golang-boilerpalte/api/v1"
)

func AddV1Router(r *gin.Engine) {
	v1Routes := r.Group("/api/v1")
	v1.AddCityRouter(v1Routes)
}
