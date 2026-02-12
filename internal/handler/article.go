package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gonews/internal/helper"
	"github.com/muhammadidrusalawi/gonews/internal/service"
)

type ArticleRequest struct {
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ArticleResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
}

func GetAllArticlesHandler(c *fiber.Ctx) error {
	articles, err := service.GetAllArticles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(helper.ApiError("Failed to fetch articles"))
	}

	var resp []ArticleResponse
	for _, a := range articles {
		resp = append(resp, ArticleResponse{
			ID:          a.ID,
			Title:       a.Title,
			Category:    a.Category,
			Description: a.Description,
			Image:       a.Image,
			UserID:      a.CreatedBy,
			Username:    a.User.Username,
		})
	}

	return c.Status(fiber.StatusOK).
		JSON(helper.ApiSuccess("Articles fetched successfully", resp))
}

func CreateArticleHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(helper.ApiError("Unauthorized"))
	}

	req := new(ArticleRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError("Invalid request"))
	}

	article, err := service.CreateArticle(
		userID.(uint),
		req.Title,
		req.Category,
		req.Description,
		req.Image,
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err.Error()))
	}

	resp := ArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		Category:    article.Category,
		Description: article.Description,
		Image:       article.Image,
		UserID:      article.CreatedBy,
	}

	return c.Status(fiber.StatusCreated).JSON(helper.ApiSuccess("Article created successfully", resp))
}

func GetMyArticlesHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(helper.ApiError("Unauthorized"))
	}

	articles, err := service.GetArticlesByOwner(userID.(uint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(helper.ApiError("Failed to fetch articles"))
	}

	var resp []ArticleResponse
	for _, a := range articles {
		resp = append(resp, ArticleResponse{
			ID:          a.ID,
			Title:       a.Title,
			Category:    a.Category,
			Description: a.Description,
			Image:       a.Image,
			UserID:      a.CreatedBy,
			Username:    a.User.Username,
		})
	}

	return c.Status(fiber.StatusOK).
		JSON(helper.ApiSuccess("My articles fetched successfully", resp))
}
