package converter

import (
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/model/repository/entities"
)

func ConvertDomainToEntity(userDomain model.UserDomainInterface) *entities.UserEntityStruct {
	return &entities.UserEntityStruct{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Name:     userDomain.GetName(),
		Age:      userDomain.GetAge(),
	}
}

func ConvertEntityToDomain(userEntity entities.UserEntityStruct) model.UserDomainInterface {
	domain := model.NewUserDomain(
		userEntity.Email,
		userEntity.Password,
		userEntity.Name,
		userEntity.Age,
	)
	domain.SetID(userEntity.ID.Hex())
	return domain
}
