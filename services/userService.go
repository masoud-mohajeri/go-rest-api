package services

import (
	"github.com/google/uuid"
	"github.com/mohajerimasoud/go-rest-api/dtos"
	"github.com/mohajerimasoud/go-rest-api/models"
	"github.com/mohajerimasoud/go-rest-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser(user *models.User, repository *repositories.UsersRepository) dtos.Response {
	userId, err := uuid.NewRandom()
	if err != nil {
		return dtos.Response{Success: false, Message: err.Error()}
	}
	user.ID = userId.String()

	// TODO : make an auto generate salt at compile time and read it here from env vars
	salt := "salt"
	hashedPass, hashingError := bcrypt.GenerateFromPassword([]byte(user.Password+salt), 1)
	if hashingError != nil {
		return dtos.Response{Success: false, Message: hashingError.Error()}
	}
	user.Password = string(hashedPass)

	result := repository.SaveUser(user)
	if result.Error != nil {
		return dtos.Response{Success: false, Message: result.Error.Error()}
	}

	return dtos.Response{Success: true, Data: user}
}
