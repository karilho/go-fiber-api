package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/controller/userController"
)

func InitRoutes(app *fiber.App, userController userController.UserControllerInterface) {

	//Eles recebem o path e um parametro (...), que no caso você pode integrar com um context
	// Ou com middlewares (JWT), vários, quantos quiser.
	//User CRUD Routes
	app.Post("/createUser", userController.CreateUser)
	app.Get("/getUserById/:userId", userController.FindUserById)
	app.Get("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	app.Put("/:userId", userController.UpdateUser)
	app.Delete("/:userId", userController.DeleteUser)

	//Login Route
	app.Post("/login", userController.LoginUser)
}
