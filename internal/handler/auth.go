package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gonews/internal/helper"
	"github.com/muhammadidrusalawi/gonews/internal/service"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}

func Register(c *fiber.Ctx) error {
	req := new(AuthRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid request"))
	}

	user, err := service.CreateUser(req.Username, req.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err.Error()))
	}

	resp := AuthResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User registered successfully", resp))
}

func Login(c *fiber.Ctx) error {
	req := new(AuthRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid request"))
	}

	user, err := service.LoginUser(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err.Error()))
	}

	resp := AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    "jwt-secret",
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User logged in successfully", resp))
}
