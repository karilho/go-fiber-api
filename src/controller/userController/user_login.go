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

func (uc *userControllerInterface) LoginUser(ctx *fiber.Ctx) error {
	logger.Info("Starting LOGIN of user VIA CONTROLLER",
		zap.String("journey", "loginUser"),
	)
	var userLoginRequest dtos.UserLoginRequest

	if err := ctx.BodyParser(&userLoginRequest); err != nil {
		logger.Error("Error parsing body: ", err,
			zap.String("journey", "createUser"))
		errRest := rest_errors.NewBadRequestError("Incorrect field error " + err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(errRest)
	}

	err := validation.ValidateStruct(userLoginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	domain := model.NewUserLoginDomain(userLoginRequest.Email, userLoginRequest.Password)

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	logger.Info("Login Sucessfully via CONTROLLER",
		zap.String("userName", domain.GetName()),
		zap.String("journey", "userLogin"),
	)

	//When the user makes sucessfully login, we will return the token in the header
	ctx.Set("Authorization", token)
	//Get Token from header to acess
	logger.Info("Token: ", zap.String("token", token))
	return ctx.Status(fiber.StatusCreated).JSON(view.ConvertDomainToResponse(domainResult))
}
