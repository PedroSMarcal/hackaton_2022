package models

import (
	"time"

	"gorm.io/gorm"
)

type Association struct {
	ID           uint `gorm:"primaryKey"`
	CNPJ         string
	SocialReason string
	User         string
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
