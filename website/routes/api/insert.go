package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/database"
)

func Insert(c *gin.Context) {
	var post database.URL
	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error Binding JSON!",
		})

		return
	}

	if len(post.URL) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "URL was not Provided!",
		})

		return
	}

	_, err = http.Get(post.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error Pinging URL!",
		})

		return
	}

	_, err = database.Insert(&database.URL{
		URL: post.URL,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error Inserting to Database!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"url":     post.URL,
	})
}
