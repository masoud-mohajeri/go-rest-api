package services

import (
	"github.com/mohajerimasoud/go-rest-api/dtos"
	"github.com/mohajerimasoud/go-rest-api/models"
	"github.com/mohajerimasoud/go-rest-api/repositories"
)

func FindAllArticles(repository repositories.ArticleRepository) dtos.Response {
	queryResult := repository.FindAll()

	if queryResult.Error != nil {
		return dtos.Response{Success: false, Message: queryResult.Error.Error()}
	}
	data := queryResult.Result.(*models.Articles)
	return dtos.Response{Success: true, Data: data}

}

func SaveArticle(article *models.Article, repository repositories.ArticleRepository) dtos.Response {
	queryResult := repository.Save(article)
	if queryResult.Error != nil {
		return dtos.Response{Success: false, Message: queryResult.Error.Error()}
	}
	// what is in result ???
	return dtos.Response{Success: true, Data: queryResult.Result}
}
