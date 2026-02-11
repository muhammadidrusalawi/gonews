package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammadidrusalawi/gonews/internal/helper"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiError("Missing Authorization header"))
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiError("Invalid Authorization header format"))
	}

	tokenStr := parts[1]
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiError("Invalid or expired token"))
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if rawID, ok := claims["user_id"].(float64); ok {
			c.Locals("user_id", uint(rawID))
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiError("Invalid user_id in token"))
		}

		if username, ok := claims["username"].(string); ok {
			c.Locals("username", username)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiError("Invalid username in token"))
		}
	}

	return c.Next()
}
