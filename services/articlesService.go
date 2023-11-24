package services

import (
	"github.com/google/uuid"
	"github.com/mohajerimasoud/go-rest-api/dtos"
	"github.com/mohajerimasoud/go-rest-api/models"
	"github.com/mohajerimasoud/go-rest-api/repositories"
	"log"
)

func FindAllArticles(repository repositories.ArticleRepository) dtos.Response {
	queryResult := repository.FindAll()

	if queryResult.Error != nil {
		return dtos.Response{Success: false, Message: queryResult.Error.Error()}
	}
	data := queryResult.Result.(models.Articles)
	return dtos.Response{Success: true, Data: data}

}

func SaveArticle(article *models.Article, repository repositories.ArticleRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}
	article.ID = uuidResult.String()

	queryResult := repository.Save(article)
	if queryResult.Error != nil {
		return dtos.Response{Success: false, Message: queryResult.Error.Error()}
	}
	// what is in result ???
	return dtos.Response{Success: true, Data: queryResult.Result}
}

func DeleteArticle(id *string, repository repositories.ArticleRepository) dtos.Response {
	queryResult := repository.DeleteArticleById(*id)

	if queryResult.Error != nil {
		return dtos.Response{Success: false, Message: queryResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: queryResult.Result}
}
