package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gonews/internal/helper"
)

func ErrorMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[PANIC RECOVER] %v\n", r)
			c.Status(fiber.StatusInternalServerError).JSON(helper.ApiError("Internal Server Error"))
		}
	}()

	err := c.Next()

	if err != nil {
		log.Printf("[ERROR] %v\n", err)

		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return c.Status(code).JSON(helper.ApiError(err.Error()))
	}

	return nil
}
