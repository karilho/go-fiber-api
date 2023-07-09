package service

import (
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/model/repository"
)

func NewUserDomainService(repository repository.UserRepositoryInterface) UserDomainService {
	return &userDomainService{
		repository,
	}
}

type userDomainService struct {
	repository repository.UserRepositoryInterface
}

// Aqui eu crio uma interface para que o controller ou quem precise possa chamar o metodo
// Relacionado com o padrão PROTOTYPE
type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_errors.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_errors.RestErr
	DeleteUser(string) *rest_errors.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_errors.RestErr)
}
