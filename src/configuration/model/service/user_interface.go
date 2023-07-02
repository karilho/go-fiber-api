package service

import (
	"github.com/karilho/go-fiber-api/src/configuration/model"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

// Aqui eu crio uma interface para que o controller ou quem precise possa chamar o metodo
// Relacionado com o padrão PROTOTYPE
type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_errors.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_errors.RestErr
	DeleteUser(string) *rest_errors.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_errors.RestErr)
}
