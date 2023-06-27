package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/controller/dtos"
)

// FiberValidator is a struct that holds the validator and translator
type ValidateUser struct {
	validator *validator.Validate
}

func NewValidateUser() *ValidateUser {
	return &ValidateUser{
		validator: validator.New(),
	}
}

func (vu *ValidateUser) Middleware(ctx *fiber.Ctx) error {
	user := new(dtos.UserRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	err := vu.validator.Struct(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Adicione o objeto User ao contexto para que o controlador possa acessá-lo
	ctx.Locals("user", user)

	// Chame o próximo manipulador na cadeia
	return ctx.Next()
}
