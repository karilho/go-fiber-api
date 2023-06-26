package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/controller/UserController"
)

func InitRoutes(app *fiber.App) {

	//Estudos -> Este ponto é para gerenciar as rotas.
	//Eles recebem o path e um parametro (...), que no caso você pode integrar com um context
	// Ou com middlewares (JWT), vários, quantos quiser.
	app.Post("/:userId", UserController.CreateUser)
	app.Get("/getUserById/:userId", UserController.FindUserById)
	app.Get("/getUserByEmail/:userEmail", UserController.FindUserByEmail)
	app.Get("/getUsers", UserController.FindUsers)
	app.Put("/:userId", UserController.UpdateUser)
	app.Delete("/:userId", UserController.DeleteUser)
	
}
