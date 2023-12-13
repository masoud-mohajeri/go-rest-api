package config

import (
	"github.com/gin-gonic/gin"
	"github.com/mohajerimasoud/go-rest-api/models"
	"github.com/mohajerimasoud/go-rest-api/repositories"
	"github.com/mohajerimasoud/go-rest-api/services"
	"net/http"
)

func UsersRoutes(router *gin.RouterGroup, UsersRepository *repositories.UsersRepository) {

	router.POST("/", func(c *gin.Context) {
		user := models.User{}
		err := c.ShouldBindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad payload")
			return
		}

		if user.Role == models.Admin {
			c.JSON(http.StatusBadRequest, "User must be promoted to be an admin")
			return
		}

		result := services.SaveUser(&user, UsersRepository)
		if !result.Success {
			c.JSON(http.StatusInternalServerError, result.Message)
			return
		}

		c.JSON(http.StatusCreated, "User created successfully")

	})

}
