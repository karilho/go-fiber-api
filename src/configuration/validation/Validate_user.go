package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
)

// Usar o translate para mudar a forma de exibição do erro para o usuário
// When you use var at the package level, the variable is initialized to its zero value, like dependency injection in Java.
var (
	validate = validator.New()
)

func ValidateStruct(request any) *rest_errors.RestErr {
	err := validate.Struct(request)
	if err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}
	return nil
}
