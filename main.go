package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/controller/routes"
	"github.com/karilho/go-fiber-api/src/controller/userController"
	"github.com/karilho/go-fiber-api/src/database/mongodb"
	"github.com/karilho/go-fiber-api/src/model/repository"
	"github.com/karilho/go-fiber-api/src/model/service"
	"log"
)

func main() {
	ctx := context.Background()
	logger.Info("Starting the application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//init db
	db, err := mongodb.NewMongoConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// init dependencies
	repository := repository.NewUserRepository(db)
	service := service.NewUserDomainService(repository)
	userController := userController.NewUserControllerInterface(service)

	app := fiber.New()
	routes.InitRoutes(app, userController)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

}
