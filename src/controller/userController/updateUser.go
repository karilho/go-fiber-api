package userController

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/controller/dtos"
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(ctx *fiber.Ctx) error {
	logger.Info("Starting update of user VIA CONTROLLER",
		zap.String("journey", "updateUser"),
	)
	var userUpdateRequest dtos.UserUpdateRequest
	var validation = validator.New()
	if err := ctx.BodyParser(&userUpdateRequest); err != nil {
		logger.Error("Error parsing body: ", err,
			zap.String("journey", "updateUser"))
		//Erro que vai retornar pro usuário no put.
		errRest := validation.Struct(userUpdateRequest)
		//Este retorno serve para que ele NÃO CONTINUE E CÓDIGO CASO ERRO
		return ctx.Status(fiber.StatusBadRequest).JSON(errRest)
	}

	userId := ctx.Params("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_errors.NewBadRequestError("Invalid ID, must be HEX value")
		ctx.Status(fiber.StatusBadRequest).JSON(errRest)
	}

	domain := model.NewUserUpdateDomain(userUpdateRequest.Name, userUpdateRequest.Age)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	logger.Info("Updated user sucessfully via CONTROLLER",
		zap.String("userID", domain.GetId()),
		zap.String("Name", domain.GetName()),
		zap.Int("Age", domain.GetAge()),
		zap.String("journey", "updateUser"),
	)

	return ctx.Status(fiber.StatusOK).JSON(view.ConvertDomainToResponse(domain))
}
