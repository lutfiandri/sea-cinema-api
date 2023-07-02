package main

import (
	"log"

	"sea-cinema-api/internal/config"
	"sea-cinema-api/internal/controller"
	"sea-cinema-api/internal/infrastructure"
	"sea-cinema-api/internal/middleware"
	"sea-cinema-api/internal/repository"
	"sea-cinema-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	appConfig := fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	}
	app := fiber.New(appConfig)
	app.Use(recover.New())

	db := infrastructure.NewMongoDatabase(config.MONGO_URI, config.MONGO_DB_NAME)

	userRepository := repository.NewUserRepository(db, "users")
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(app, authService)
	authController.InitRoute()

	log.Fatal(app.Listen(":4000"))
}
