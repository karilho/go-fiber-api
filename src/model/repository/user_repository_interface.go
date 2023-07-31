package repository

import (
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_COLLECTION = "MONGO_DB_COLLECTION"
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
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_errors.RestErr)
	FindUserById(id string) (model.UserDomainInterface, *rest_errors.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_errors.RestErr
	DeleteUser(userId string) *rest_errors.RestErr
	FindUserByEmailAndPass(email string, password string) (model.UserDomainInterface, *rest_errors.RestErr)
}
