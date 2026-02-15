package handler

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/muhammadidrusalawi/gonews/internal/helper"
)

func UploadImageHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError("Image is required"))
	}

	ext := filepath.Ext(file.Filename)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp":
	default:
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError("Invalid image format"))
	}

	filename := uuid.New().String() + ext
	path := "./public/uploads/" + filename

	if err := c.SaveFile(file, path); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(helper.ApiError("Failed to save image"))
	}

	imageURL := "/uploads/" + filename

	return c.Status(fiber.StatusCreated).
		JSON(helper.ApiSuccess("Image uploaded successfully", fiber.Map{
			"image_url": imageURL,
		}))
}
