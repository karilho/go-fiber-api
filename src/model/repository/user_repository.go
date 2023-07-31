package repository

import (
	"context"
	"fmt"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/model/repository/entities"
	"github.com/karilho/go-fiber-api/src/model/repository/entities/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
)

func (urs *userRepositoryStruct) CreateUser(userDomain model.UserDomainInterface) (
	model.UserDomainInterface, *rest_errors.RestErr) {

	logger.Info("Starting creation of user VIA REPOSITORY")

	COLLECTION_NAME := os.Getenv(MONGO_COLLECTION)

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

func (urs *userRepositoryStruct) FindUserByEmail(email string) (model.UserDomainInterface, *rest_errors.RestErr) {

	logger.Info("Starting GETTER EMAIL of user VIA REPOSITORY")

	COLLECTION_NAME := os.Getenv(MONGO_COLLECTION)
	collection := urs.databaseConnection.Collection(COLLECTION_NAME)
	userEntity := &entities.UserEntityStruct{}
	filter := bson.D{{"email", email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMsg := fmt.Sprintf(("User with email %s not found"), email)
			logger.Error(errorMsg, err,
				zap.String("journey", "FindUserByEmail"))
			return nil, rest_errors.NewNotFoundError(errorMsg)
		}

		errorMsg := "Error trying to find  user by email"
		return nil, rest_errors.NewInternalServerError(errorMsg)
	}
	logger.Info("User found",
		zap.String("journey", "FindUserByEmail"),
		zap.String("email", email),
		zap.String("journey", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (urs *userRepositoryStruct) FindUserById(id string) (model.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("Starting GETTER ID of user VIA REPOSITORY")

	COLLECTION_NAME := os.Getenv(MONGO_COLLECTION)
	collection := urs.databaseConnection.Collection(COLLECTION_NAME)
	userEntity := &entities.UserEntityStruct{}
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMsg := fmt.Sprintf(("User with id %s not found"), id)
			logger.Error(errorMsg, err,
				zap.String("journey", "findUserById"))
			return nil, rest_errors.NewNotFoundError(errorMsg)
		}

		errorMsg := "Error trying to find  user by ID"
		return nil, rest_errors.NewInternalServerError(errorMsg)
	}
	logger.Info("User found",
		zap.String("journey", "findUserById"),
		zap.String("journey", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (urs *userRepositoryStruct) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_errors.RestErr {

	logger.Info("Starting update of user VIA REPOSITORY")

	COLLECTION_NAME := os.Getenv(MONGO_COLLECTION)
	collection := urs.databaseConnection.Collection(COLLECTION_NAME)

	value := converter.ConvertDomainToEntity(userDomain)
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{"_id", userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user",
			err,
			zap.String("journey", "updateUser"))
		return rest_errors.NewInternalServerError(err.Error())
	}

	logger.Info("User updated",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	return nil
}

func (urs *userRepositoryStruct) DeleteUser(userId string) *rest_errors.RestErr {
	logger.Info("Starting delete of user VIA REPOSITORY",
		zap.String("journey", "deleteUser"))

	COLLECTION_NAME := os.Getenv(MONGO_COLLECTION)
	collection := urs.databaseConnection.Collection(COLLECTION_NAME)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{"_id", userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user",
			err,
			zap.String("journey", "deleteUser"))
		return rest_errors.NewInternalServerError(err.Error())
	}

	logger.Info("User deleted sucessfully via Repository",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	return nil
}
