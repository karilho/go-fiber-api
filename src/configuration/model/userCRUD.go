package model

import (
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_errors.RestErr {
	logger.Info("Starting creation of user VIA MODEL",
		zap.String("journey", "CreateUser"))
	ud.EncryptPass()
	logger.Info("User ALIAS",
		zap.String("userName", ud.Name),
		zap.String("EncryptPass", ud.Password))
	return nil
}

func (ud *UserDomain) UpdateUser(string) *rest_errors.RestErr {
	return nil
}

func (ud *UserDomain) DeleteUser(string) *rest_errors.RestErr {
	return nil
}

func (ud *UserDomain) FindUser(string) (*UserDomain, *rest_errors.RestErr) {
	return nil, nil
}
