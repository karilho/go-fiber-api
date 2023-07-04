package view

import (
	"github.com/karilho/go-fiber-api/src/controller/dtos"
	"github.com/karilho/go-fiber-api/src/model"
)

func ConvertDomainToResponse(userDomainInterface model.UserDomainInterface) dtos.UserResponse {
	return dtos.UserResponse{
		//ID:    userDomainInterface.GetID(),
		Email: userDomainInterface.GetEmail(),
		Name:  userDomainInterface.GetName(),
		Age:   userDomainInterface.GetAge(),
	}
}
