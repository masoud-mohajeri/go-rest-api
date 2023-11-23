package config

import (
	"github.com/gin-gonic/gin"
	"github.com/mohajerimasoud/go-rest-api/helpers"
	"github.com/mohajerimasoud/go-rest-api/models"
	"github.com/mohajerimasoud/go-rest-api/repositories"
	"github.com/mohajerimasoud/go-rest-api/services"
	"net/http"
)

func setupRoutes(ArticlesRepository *repositories.ArticleRepository) {

	route := gin.Default()

	route.GET("/articles", func(context *gin.Context) {
		code := http.StatusOK
		response := services.FindAllArticles(*ArticlesRepository)

		if !response.Success {
			code = http.StatusNotFound
		}

		context.JSON(code, response)
	})

	route.POST("/articles", func(context *gin.Context) {
		var article models.Article
		err := context.ShouldBindJSON(&article)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			context.JSON(http.StatusBadRequest, response)
			return
		}

		response := services.SaveArticle(&article, *ArticlesRepository)
		code := http.StatusCreated

		if !response.Success {
			code = http.StatusBadRequest
			context.JSON(code, response)
			return
		}

		context.JSON(code, response)
	})

}
