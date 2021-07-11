package api

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/url", GetAll)
		api.POST("/url", Insert)
		api.DELETE("/url/:url", Delete)

		api.GET("/status", Status)
	}
}
