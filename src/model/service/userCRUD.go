package service

import (
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) CreateUser(ud model.UserDomainInterface) *rest_errors.RestErr {
	logger.Info("Starting creation of user VIA MODEL",
		zap.String("journey", "CreateUser"))
	ud.EncryptPass()
	logger.Info("User ALIAS",
		zap.String("userName", ud.GetName()),
		zap.String("EncryptPass", ud.GetPassword()))
	return nil
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
