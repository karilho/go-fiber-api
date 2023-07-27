package service

import (
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/model/repository"
)

// Quando eu chamar essa classse como por exemplo no crud, como ela ta chamando o repository eu preciso passar o repository
// Aqui eu crio uma interface para que o controller ou quem precise possa chamar o metodo
// Relacionado com o padrão PROTOTYPE
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
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_errors.RestErr
	DeleteUser(string) *rest_errors.RestErr
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_errors.RestErr)
	FindUserById(id string) (model.UserDomainInterface, *rest_errors.RestErr)
}
