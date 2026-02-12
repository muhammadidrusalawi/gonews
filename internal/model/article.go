package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(200);not null"`
	Category    string `gorm:"type:varchar(100);not null;index"`
	Description string `gorm:"type:text;not null"`
	Image       string `gorm:"type:varchar(255);not null"`
	CreatedBy   uint   `gorm:"index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:CreatedBy"`
}
