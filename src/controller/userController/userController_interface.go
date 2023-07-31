package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/model/service"
)

func NewUserControllerInterface(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	FindUserById(c *fiber.Ctx) error
	FindUserByEmail(c *fiber.Ctx) error

	UpdateUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
}

type userControllerInterface struct {
	service service.UserDomainService
}
