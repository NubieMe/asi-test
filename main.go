package main

import (
	"log"
	"test-asi/handler"
	"test-asi/internal/config"
	"test-asi/repository"
	"test-asi/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	redisClient, redisStore := config.ConnectRedis()

	userRepo := repository.NewUserRepository(redisClient)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService, redisStore)

	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	log.Fatal(app.Listen(":3000"))
}
