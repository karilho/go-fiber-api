package service_tests

import (
	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/model/service"
	mock "github.com/karilho/go-fiber-api/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
)

func TestUserDomainService_LoginUserServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepositoryInterface(ctrl)
	service := service.NewUserDomainService(repository)

	t.Run("when_calling_repository_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		userDomainMock := model.NewUserDomain(
			userDomain.GetEmail(),
			userDomain.GetPassword(),
			userDomain.GetName(),
			userDomain.GetAge())
		userDomainMock.EncryptPass()

		repository.EXPECT().FindUserByEmailAndPass(
			userDomain.GetEmail(), userDomainMock.GetPassword()).Return(
			nil, rest_errors.NewInternalServerError("error trying to find user by email and password"))

		user, token, err := service.LoginUserService(userDomain)
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to find user by email and password")
	})

	t.Run("when_calling_create_token_returns_error", func(t *testing.T) {
		userDomainMock := mock.NewMockUserDomainInterface(ctrl)

		userDomainMock.EXPECT().GetEmail().Return("test@test.com")
		userDomainMock.EXPECT().GetPassword().Return("test")
		userDomainMock.EXPECT().EncryptPass()

		userDomainMock.EXPECT().GenerateToken().Return("",
			rest_errors.NewInternalServerError("error trying to create token"))

		repository.EXPECT().FindUserByEmailAndPass(
			"test@test.com", "test").Return(
			userDomainMock, nil)

		user, token, err := service.LoginUserService(userDomainMock)
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create token")
	})

	t.Run("when_user_and_password_is_valid_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		secret := "secret"

		err := os.Setenv("JWT_SECRET_KEY", secret)
		if err != nil {
			t.FailNow()
			return
		}
		defer os.Clearenv()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPass(
			userDomain.GetEmail(), gomock.Any()).Return(
			userDomain, nil)

		userDomainReturn, token, err := service.LoginUserService(userDomain)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetId(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

		tokenReturned, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}

			return nil, rest_errors.NewBadRequestError("invalid token")
		})
		_, ok := tokenReturned.Claims.(jwt.MapClaims)
		if !ok || !tokenReturned.Valid {
			return
		}
	})
}
