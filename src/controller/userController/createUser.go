package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/configuration/validation"
	"github.com/karilho/go-fiber-api/src/controller/dtos"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateUser(ctx *fiber.Ctx) error {
	logger.Info("Starting creation of user",
		zap.String("journey", "createUser"),
	)
	var userRequest dtos.UserRequest
	//var userResponse dtos.UserResponse
	//Se ele executar, e DER UM ERRO DIFERENTE DE RETORNO NULO
	// ELE VAI LANÇAR ESSA EXCEÇÃO NAS LINHAS 17/18
	// SE DER TUDO BEM, VAI FAZER O BODYPARSER
	if err := ctx.BodyParser(&userRequest); err != nil {
		logger.Error("Error parsing body: ", err,
			zap.String("journey", "createUser"))
		//Erro que vai retornar pro usuário no post.
		errRest := rest_errors.NewBadRequestError("Incorrect field error " + err.Error())
		//Este retorno serve para que ele NÃO CONTINUE E CÓDIGO CASO ERRO
		return ctx.Status(fiber.StatusBadRequest).JSON(errRest)
	}

	//Aqui eu vou validar o usuário
	err := validation.ValidateStruct(userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	//Se tudo correr bem, retorne o response, mas ai tenho q ver kkk
	userResponse := dtos.UserResponse{
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("Creater user sucessfully",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})
	return ctx.JSON(userResponse)
}
