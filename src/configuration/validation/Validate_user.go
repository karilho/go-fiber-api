package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/controller/dtos"
)

// Usar o translate para mudar a forma de exibição do erro para o usuário
// Quando vocÊ chama no var no começo, é como se fosse injetar (java) o que você irá utilizar.
var (
	validate = validator.New()
)

func ValidateStruct(user dtos.UserRequest) *rest_errors.RestErr {
	err := validate.Struct(user)
	if err != nil {
		var causes []rest_errors.Causes
		for _, err := range err.(validator.ValidationErrors) {
			cause := rest_errors.Causes{
				Message: err.Error(),
				Field:   err.Field(),
			}
			causes = append(causes, cause)

		}
		restErr := rest_errors.NewBadRequestValidationError("Invalid request body", causes)
		return restErr
	}
	return nil
}
