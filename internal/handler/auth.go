package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gonews/internal/config"
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

	token, err := config.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helper.ApiError("Failed to generate token"))
	}

	resp := AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
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

	token, err := config.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helper.ApiError("Failed to generate token"))
	}

	resp := AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User logged in successfully", resp))
}

func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiError("Unauthorized"))
	}

	user, err := service.GetUser(userID.(uint))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(helper.ApiError(err.Error()))
	}

	resp := AuthResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User retrieved successfully", resp))
}
