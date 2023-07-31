package service

import (
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_errors.RestErr) {
	logger.Info("Starting LOGIN of user VIA SERVICE",
		zap.String("journey", "loginUser"))

	userDomain.EncryptPass()

	user, err := uds.repository.FindUserByEmailAndPass(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	tokenJWT, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info("User found via SERVICE",
		zap.String("userID", user.GetId()),
		zap.String("journey", "loginUser"))

	return user, tokenJWT, nil
}
