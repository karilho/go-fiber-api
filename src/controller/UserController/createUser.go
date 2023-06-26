package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
)

func CreateUser(c *fiber.Ctx) {
	err := rest_errors.NewBadRequestError("This is a bad request")
	c.JSON(err)

}
