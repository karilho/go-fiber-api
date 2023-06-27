package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/controller/dtos"
)

func CreateUser(ctx *fiber.Ctx) error {

	// Acesse o objeto User do contexto
	user := ctx.Locals("user").(*dtos.UserRequest)

	// Lógica para criar o usuário

	return ctx.Status(fiber.StatusCreated).JSON(user)
}
