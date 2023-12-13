package config

import (
	"github.com/gin-gonic/gin"
	"github.com/mohajerimasoud/go-rest-api/repositories"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	ArticlesRouter := router.Group("/articles")
	ArticlesRepository := repositories.NewArticleRepository(db)
	ArticlesRoutes(ArticlesRouter, ArticlesRepository)

	return router

}
