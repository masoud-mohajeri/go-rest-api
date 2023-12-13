package repositories

import (
	"github.com/mohajerimasoud/go-rest-api/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

// Is it nessasary???
func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) SaveUser(user *models.User) RepositoryResult {
	err := u.db.Create(user).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: user}

}

func (u *UsersRepository) ChangeUserRole(id *string, ut models.UserType) RepositoryResult {
	user := &models.User{}
	user.ID = *id
	user.Role = ut
	tx := u.db.Updates(user)
	if tx.Error != nil {
		return RepositoryResult{Error: tx.Error}
	}

	return RepositoryResult{Result: user}
}
