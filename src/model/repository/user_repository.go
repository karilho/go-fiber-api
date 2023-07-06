package repository

import (
	"context"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"os"
)

func (urs *userRepositoryStruct) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_errors.RestErr) {
	//urs.databaseConnection.
	logger.Info("Starting creation of user VIA REPOSITORY")

	COLLECTION_NAME := os.Getenv("MONGO_DB_COLLECTION")
	collection := urs.databaseConnection.Collection(COLLECTION_NAME)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		logger.Error("Error on get JSON value", err)
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error on insert user", err)
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))
	return userDomain, nil
}
