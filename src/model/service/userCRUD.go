package service

import (
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) CreateUser(udi model.UserDomainInterface) (model.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("Starting creation of user VIA MODEL -> Service layer",
		zap.String("journey", "CreateUser"))
	udi.EncryptPass()

	userDomainRepository, err := uds.repository.CreateUser(udi)
	if err != nil {
		logger.Error("Error on insert user",
			err,
			zap.String("error", err.Error()))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	logger.Info("User ALIAS",
		zap.String("userName", udi.GetName()),
		zap.String("EncryptPass", udi.GetPassword()))
	return userDomainRepository, nil
}

func (ud *userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_errors.RestErr {
	logger.Info("Starting update of user VIA MODEL -> Service layer",
		zap.String("journey", "updateUser"))

	err := ud.repository.UpdateUser(userID, userDomain)
	if err != nil {
		logger.Error("Error on update user when calling repository", err)
		return err
	}

	logger.Info("User ALIAS",
		zap.String("Id", userID),
		zap.String("Name", userDomain.GetName()))
	return nil

}

func (urs *userDomainService) DeleteUser(userId string) *rest_errors.RestErr {
	logger.Info("Starting delete of user VIA MODEL -> Service layer",
		zap.String("journey", "deleteUser"))

	err := urs.repository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error on delete user when calling repository",
			err,
			zap.String("error", err.Error()))
	}

	logger.Info("User deleted sucessfully via MODEL -> Service layer",
		zap.String("userID", userId),
		zap.String("journey", "deleteUser"))
	return err
}

func (urs *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("Starting GETTER EMAIL of user VIA SERVICE -> Service layer",
		zap.String("journey", "FindUserByEmail"))

	// Pq se eu retornar direto não dá?
	userDomain, err := urs.repository.FindUserByEmail(email)
	if err != nil {
		return nil, err

	}
	return userDomain, nil
}

func (urs *userDomainService) FindUserById(id string) (model.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("Starting GETTER ID of user VIA SERVICE -> Service layer",
		zap.String("journey", "FindUserById"))

	// Pq se eu retornar direto não dá?
	userDomain, err := urs.repository.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return userDomain, nil
}
