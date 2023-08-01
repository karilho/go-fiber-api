package service

import (
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/model/repository"
)

// When i call this function, i will pass the repository that i want to use like a dependency injection
func NewUserDomainService(repository repository.UserRepositoryInterface) UserDomainService {
	return &userDomainService{
		repository,
	}
}

type userDomainService struct {
	repository repository.UserRepositoryInterface
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_errors.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_errors.RestErr
	DeleteUser(string) *rest_errors.RestErr
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_errors.RestErr)
	FindUserById(id string) (model.UserDomainInterface, *rest_errors.RestErr)
	LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_errors.RestErr)
}
