package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/controller/userController"
)

func InitRoutes(app *fiber.App, userController userController.UserControllerInterface) {

	//Estudos -> Este ponto é para gerenciar as rotas.
	//Eles recebem o path e um parametro (...), que no caso você pode integrar com um context
	// Ou com middlewares (JWT), vários, quantos quiser.
	app.Post("/createUser", userController.CreateUser)
	app.Get("/getUserById/:userId", userController.FindUserById)
	app.Get("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	app.Get("/getUsers", userController.FindUsers)
	app.Put("/:userId", userController.UpdateUser)
	app.Delete("/:userId", userController.DeleteUser)

}
