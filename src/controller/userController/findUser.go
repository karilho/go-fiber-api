package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *fiber.Ctx) error {
	logger.Info("Starting GETTER EMAIL of user VIA CONTROLLER",
		zap.String("journey", "FindUserById"))

	userId := c.Params("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMsg := rest_errors.NewBadRequestError(
			"Invalid ID")
		return c.Status(errorMsg.Code).JSON(errorMsg)
	}

	userDomain, err := uc.service.FindUserById(userId)
	if err != nil {
		return c.Status(err.Code).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *fiber.Ctx) error {
	logger.Info("Starting GETTER EMAIL of user VIA CONTROLLER",
		zap.String("journey", "FindUserByEmail"))

	userEmail := c.Params("userEmail")

	/*
		if _, err := uuid.Parse(userEmail); err != nil {
			errorMsg := rest_errors.NewBadRequestError(
				"Invalid email")
			return c.Status(errorMsg.Code).JSON(errorMsg)
		}

	*/

	userDomain, err := uc.service.FindUserByEmail(userEmail)
	if err != nil {
		return c.Status(err.Code).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUsers(c *fiber.Ctx) error {
	return nil
}
