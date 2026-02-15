package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/muhammadidrusalawi/gonews/internal/database"
	"github.com/muhammadidrusalawi/gonews/internal/middleware"
	"github.com/muhammadidrusalawi/gonews/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	database.InitDB()

	app := fiber.New()
	app.Use(middleware.ErrorMiddleware)
	app.Use(logger.New())

	app.Static("/", "./public")
	app.Static("/uploads", "./public/uploads")
	routes.ApiRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
