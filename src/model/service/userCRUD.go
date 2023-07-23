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

func (*userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_errors.RestErr {
	return nil
}

func (*userDomainService) DeleteUser(string) *rest_errors.RestErr {
	return nil
}

func (*userDomainService) FindUser(string) (*model.UserDomainInterface, *rest_errors.RestErr) {
	return nil, nil
}
