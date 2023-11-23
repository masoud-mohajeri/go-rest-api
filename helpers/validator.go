package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/mohajerimasoud/go-rest-api/dtos"
	"github.com/mohajerimasoud/go-rest-api/langs"
)

// Credits to guy in reddit: @karam_v8

func GenerateValidationResponse(err error) (response dtos.ValidationResponse) {
	response.Success = false

	var validations []dtos.Validation

	// get validation errors
	validationErrors := err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		field, rule := value.Field, value.Tag
		// does it work ???
		validation := dtos.Validation{Field: field(), Message: langs.GenerateValidationMessage(field(), rule())}
		validations = append(validations, validation)
	}

	// set Validations response
	response.Validations = validations

	return response
}
