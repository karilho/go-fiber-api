package repository

import (
	"context"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/model/repository/entities/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

func (urs *userRepositoryStruct) CreateUser(userDomain model.UserDomainInterface) (
	model.UserDomainInterface, *rest_errors.RestErr) {
	//urs.databaseConnection.
	logger.Info("Starting creation of user VIA REPOSITORY")

	COLLECTION_NAME := os.Getenv("MONGO_DB_COLLECTION")

	collection := urs.databaseConnection.Collection(COLLECTION_NAME)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error on insert user", err)
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	value.ID = result.InsertedID.(primitive.ObjectID)
	return converter.ConvertEntityToDomain(*value), nil
}
