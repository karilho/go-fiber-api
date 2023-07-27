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

func ValidateStruct(userRequest dtos.UserRequest) *rest_errors.RestErr {
	err := validate.Struct(userRequest)
	if err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}
	return nil
}

func ValidateUpdate(request dtos.UserUpdateRequest) *rest_errors.RestErr {
	err := validate.Struct(request)
	if err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}
	return nil
}
