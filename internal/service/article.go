package service

import (
	"github.com/muhammadidrusalawi/gonews/internal/database"
	"github.com/muhammadidrusalawi/gonews/internal/model"
)

func GetAllArticles() ([]model.Article, error) {
	var articles []model.Article

	if err := database.DB.
		Preload("User").
		Order("created_at DESC").
		Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func CreateArticle(userID uint, title, category, description, image string) (*model.Article, error) {
	article := model.Article{
		Title:       title,
		Category:    category,
		Description: description,
		Image:       image,
		CreatedBy:   userID,
	}

	if err := database.DB.Create(&article).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

func GetArticlesByOwner(userID uint) ([]model.Article, error) {
	var articles []model.Article

	if err := database.DB.
		Preload("User").
		Where("created_by = ?", userID).
		Order("created_at DESC").
		Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}
