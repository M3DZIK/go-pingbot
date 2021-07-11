package api

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/database"
)

func GetAll(c *gin.Context) {
	results, err := database.GetAll()

	// Error Handling
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error fetching URLs!",
		})
		fmt.Println(err)

		return
	}

	// DB Is Empty
	if results == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Database is empty!",
		})

		return
	}

	var db []dbType

	for _, value := range results {
		hash := base64.URLEncoding.EncodeToString([]byte(value.URL))

		db = append(db, dbType{
			URL:  value.URL,
			HASH: hash,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"db":      db,
	})
}

type dbType struct {
	URL  string
	HASH string
}
