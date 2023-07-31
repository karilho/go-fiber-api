package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(ctx *fiber.Ctx) error {
	logger.Info("Starting delete of user VIA CONTROLLER",
		zap.String("journey", "deleteUser"),
	)

	userId := ctx.Params("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_errors.NewBadRequestError("Invalid ID, must be HEX value")
		ctx.Status(fiber.StatusBadRequest).JSON(errRest)
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	logger.Info("USER deleted sucessfully via CONTROLLER",
		zap.String("userID", userId),
		zap.String("journey", "deleteUser"),
	)

	return ctx.Status(fiber.StatusOK).JSON("User deleted sucessfully")
}
