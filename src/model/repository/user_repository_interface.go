package repository

import (
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepositoryInterface {
	return &userRepositoryStruct{
		databaseConnection: database,
	}
}

// Essa struct precisará da dependencia de conexão com DB
type userRepositoryStruct struct {
	databaseConnection *mongo.Database
}

type UserRepositoryInterface interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_errors.RestErr)
}
