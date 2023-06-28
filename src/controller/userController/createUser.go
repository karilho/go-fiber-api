package userController

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"github.com/karilho/go-fiber-api/src/controller/dtos"
	"log"
)

func CreateUser(ctx *fiber.Ctx) error {
	var userRequest dtos.UserRequest
	//var userResponse dtos.UserResponse
	//Se ele executar, e DER UM ERRO DIFERENTE DE RETORNO NULO
	// ELE VAI LANÇAR ESSA EXCEÇÃO NAS LINHAS 17/18
	// SE DER TUDO BEM, VAI FAZER O BODYPARSER
	if err := ctx.BodyParser(&userRequest); err != nil {
		log.Printf("Error parsing body: %v", err.Error())
		//Erro que vai retornar pro usuário no post.
		errRest := rest_errors.NewBadRequestError("Incorrect field error " + err.Error())
		//Este retorno serve para que ele NÃO CONTINUE E CÓDIGO CASO ERRO
		return ctx.Status(fiber.StatusBadRequest).JSON(errRest)
	}
	fmt.Println(userRequest)

	//Se tudo correr bem, retorne o response, mas ai tenho q ver kkk
	userResponse := dtos.UserResponse{
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	fmt.Println(userResponse)
	return ctx.JSON(userResponse)
}
