package website

import (
	"os"
	"strconv"

	"github.com/MedzikUser/go-utils/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/MedzikUser/go/pingbot/config"
	"gitlab.com/MedzikUser/go/pingbot/website/routes"
	"gitlab.com/MedzikUser/go/pingbot/website/routes/api"
)

var router *gin.Engine

func Server() {
	// Disable GIN Debug
	gin.SetMode(gin.ReleaseMode)

	router = gin.New()

	// Fix cors
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"https://www.pingbot.cf"}

	router.Use(cors.New(configCors))

	router.GET("/", routes.Index)

	api.ApplyRoutes(router)

	err := router.Run(":" + strconv.Itoa(config.Toml.HTTP.Port))
	if common.CheckErr(err, "gin server run") {
		os.Exit(1)
	}
}
