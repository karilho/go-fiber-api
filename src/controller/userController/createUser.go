package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/configuration/validation"
	"github.com/karilho/go-fiber-api/src/controller/dtos"
	"github.com/karilho/go-fiber-api/src/model"
	"github.com/karilho/go-fiber-api/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(ctx *fiber.Ctx) error {
	logger.Info("Starting creation of user VIA CONTROLLER",
		zap.String("journey", "createUser"),
	)
	var userRequest dtos.UserRequest

	if err := ctx.BodyParser(&userRequest); err != nil {
		logger.Error("Error parsing body: ", err,
			zap.String("journey", "createUser"))
		errRest := rest_errors.NewBadRequestError("Incorrect field error " + err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(errRest)
	}

	//TODO - Validate request inside bodyparser, not alone.
	err := validation.ValidateStruct(userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	logger.Info("Created user sucessfully via CONTROLLER",
		zap.String("userName", domain.GetName()),
		zap.String("encryptPass", domain.GetPassword()),
		zap.String("journey", "createUser"),
	)

	return ctx.Status(fiber.StatusCreated).JSON(view.ConvertDomainToResponse(domainResult))
}
