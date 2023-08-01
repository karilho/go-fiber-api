package service_tests

import (
	"github.com/golang/mock/gomock"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model/service"
	mock "github.com/karilho/go-fiber-api/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestUserDomainService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepositoryInterface(ctrl)
	service := service.NewUserDomainService(repository)

	t.Run("when_sending_a_valid_userId_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().DeleteUser(id).Return(nil)

		err := service.DeleteUser(id)

		assert.Nil(t, err)
	})

	t.Run("when_sending_a_invalid_userId_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().DeleteUser(id).Return(
			rest_errors.NewInternalServerError("error trying to update user"))

		err := service.DeleteUser(id)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
	})
}
