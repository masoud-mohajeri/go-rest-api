package config

import (
	"github.com/gin-gonic/gin"
	"github.com/mohajerimasoud/go-rest-api/helpers"
	"github.com/mohajerimasoud/go-rest-api/models"
	"github.com/mohajerimasoud/go-rest-api/repositories"
	"github.com/mohajerimasoud/go-rest-api/services"
	"net/http"
)

func SetupRoutes(ArticlesRepository *repositories.ArticleRepository) *gin.Engine {

	router := gin.Default()

	router.GET("/articles", func(context *gin.Context) {
		code := http.StatusOK
		response := services.FindAllArticles(*ArticlesRepository)

		if !response.Success {
			code = http.StatusNotFound
		}

		context.JSON(code, response)
	})

	router.POST("/articles", func(context *gin.Context) {
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

	router.DELETE("/articles/:id", func(context *gin.Context) {
		id := context.Param("id")
		code := http.StatusOK
		response := services.DeleteArticle(&id, *ArticlesRepository)
		if !response.Success {
			code = http.StatusBadRequest
			context.JSON(code, response.Message)
			return
		}
		context.JSON(code, response.Message)
	})

	router.PUT("/articles/:id", func(context *gin.Context) {
		id := context.Param("id")
		code := http.StatusOK

		var article models.Article
		err := context.ShouldBindJSON(&article)

		if err != nil {
			code = http.StatusBadRequest
			context.JSON(code, err.Error())
			return
		}

		updateResult := services.UpdateArticleById(&id, article, *ArticlesRepository)
		if !updateResult.Success {
			code = http.StatusInternalServerError
			context.JSON(code, updateResult.Message)
			return
		}

		context.JSON(code, updateResult.Message)

	})

	router.GET("/articles/:id", func(context *gin.Context) {
		id := context.Param("id")
		code := http.StatusOK

		articleSearchResult := services.FindArticleById(&id, *ArticlesRepository)

		if !articleSearchResult.Success {
			code := http.StatusNotFound
			context.JSON(code, articleSearchResult.Message)
			return
		}

		context.JSON(code, articleSearchResult.Data)
	})

	return router

}
