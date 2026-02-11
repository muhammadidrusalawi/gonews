package service

import (
	"errors"

	"github.com/muhammadidrusalawi/gonews/internal/database"
	"github.com/muhammadidrusalawi/gonews/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) (*model.User, error) {
	var existing model.User
	if err := database.DB.Where("username = ?", username).First(&existing).Error; err == nil {
		return nil, errors.New("username already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username: username,
		Password: string(hashed),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func LoginUser(username, password string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func GetUser(userID uint) (*model.User, error) {
	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
