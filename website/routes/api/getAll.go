package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
)

func GetAll(c *gin.Context) {
	results, err := mongo.GetAll()
	// Error Handling
	if err != nil {
		c.JSON(http.StatusInternalServerError, json{
			"success": false,
			"message": "Error fetching URLs!",
		})
		fmt.Println(err)

		return
	}

	// DB Is Empty
	if results == nil {
		c.JSON(http.StatusNotFound, json{
			"success": false,
			"message": "Database is empty!",
		})

		return
	}

	var db []string

	for _, value := range results {
		db = append(db, value.URL)
	}

	c.JSON(http.StatusOK, json{
		"success": true,
		"db":      db,
	})
}
