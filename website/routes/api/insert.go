package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
	"gitlab.com/gaming0skar123/go/pingbot/website/routes/api/auth"
)

func Insert(c *gin.Context) {
	valid, claims, err := auth.Authorize(c)
	if !valid {
		c.JSON(http.StatusBadRequest, json{
			"success": false,
			"message": "Unauthed!",
		})

		return
	}

	var user string

	for key, val := range claims {
		if key == "user" {
			user = fmt.Sprintf("%q", val)
			break
		}
	}

	var post mongo.URL
	err = c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, json{
			"success": false,
			"message": "Error Binding JSON!",
		})

		return
	}

	if len(post.URL) < 1 {
		c.JSON(http.StatusBadRequest, json{
			"success": false,
			"message": "URL was not Provided!",
		})

		return
	}

	_, err = http.Get(post.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, json{
			"success": false,
			"message": "Error Pinging URL!",
		})

		return
	}

	if post.Cluster == 0 {
		post.Cluster = 1
	}

	_, err = mongo.Insert(&mongo.URL{
		URL:     post.URL,
		Cluster: post.Cluster,
		Owner:   user,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, json{
			"success": false,
			"message": "Error Inserting to Database!",
		})

		return
	}

	c.JSON(http.StatusOK, json{
		"success": true,
		"url":     post.URL,
	})
}
