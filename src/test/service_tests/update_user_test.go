package service_tests

import (
	"github.com/golang/mock/gomock"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/model/service"
	mock "github.com/karilho/go-fiber-api/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestUserDomainService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepositoryInterface(ctrl)
	service := service.NewUserDomainService(repository)

	t.Run("when_sending_a_valid_user_and_userId_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(id, userDomain).Return(nil)

		err := service.UpdateUser(id, userDomain)

		assert.Nil(t, err)
	})

	t.Run("when_sending_a_invalid_user_and_userId_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(id, userDomain).Return(
			rest_errors.NewInternalServerError("error trying to update user"))

		err := service.UpdateUser(id, userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
	})
}
