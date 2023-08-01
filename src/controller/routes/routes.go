package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/controller/userController"
	"github.com/karilho/go-fiber-api/src/model"
)

func InitRoutes(app *fiber.App, userController userController.UserControllerInterface) {

	//User CRUD Routes
	app.Post("/createUser", userController.CreateUser)
	app.Get("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserById)
	app.Get("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	app.Put("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	app.Delete("/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)

	//Login Route
	app.Post("/login", userController.LoginUser)
}
