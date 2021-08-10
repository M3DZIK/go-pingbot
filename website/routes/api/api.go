package api

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/url", GetAll)

		api.GET("/status", Status)
	}
}
