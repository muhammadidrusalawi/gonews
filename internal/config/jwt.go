package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, username string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	expireStr := os.Getenv("JWT_EXPIRE")

	expire, err := time.ParseDuration(expireStr)
	if err != nil {
		expire = 24 * time.Hour
	}

	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(expire).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
