package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/muhammadidrusalawi/gonews/internal/database"
	"github.com/muhammadidrusalawi/gonews/internal/handler"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.InitDB()

	app := fiber.New()
	app.Use(logger.New())

	app.Static("/", "./public")

	api := app.Group("/api")

	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "OK!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
