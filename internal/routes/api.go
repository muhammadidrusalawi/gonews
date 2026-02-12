package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gonews/internal/handler"
	"github.com/muhammadidrusalawi/gonews/internal/helper"
	"github.com/muhammadidrusalawi/gonews/internal/middleware"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("OK!", nil))
	})

	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)
	api.Get("/profile", middleware.AuthMiddleware, handler.GetProfile)

	api.Get("/articles", handler.GetAllArticlesHandler)
	api.Post("/articles", middleware.AuthMiddleware, handler.CreateArticleHandler)
	api.Get("/my-articles", middleware.AuthMiddleware, handler.GetMyArticlesHandler)
}
