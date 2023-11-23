package config

import (
	"github.com/gin-gonic/gin"
	"github.com/mohajerimasoud/go-rest-api/repositories"
	"github.com/mohajerimasoud/go-rest-api/services"
	"net/http"
)

func setupRoutes(ArticlesRepository *repositories.ArticleRepository) {

	route := gin.Default()

	route.GET("/articles", func(c *gin.Context) {
		code := http.StatusOK
		response := services.FindAllArticles(*ArticlesRepository)

		if !response.Success {
			code = http.StatusNotFound
		}

		c.JSON(code, response)
	})

}
