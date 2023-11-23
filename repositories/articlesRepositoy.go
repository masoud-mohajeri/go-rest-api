package repositories

import (
	"github.com/mohajerimasoud/go-rest-api/models"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) Save(contact *models.Article) RepositoryResult {
	err := r.db.Save(contact).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: contact}
}

func (r *ArticleRepository) FindAll() RepositoryResult {
	var articles models.Articles
	err := r.db.Find(articles).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: articles}

}
