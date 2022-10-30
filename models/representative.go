package models

import (
	"time"

	"gorm.io/gorm"
)

type Representative struct {
	ID           uint `gorm:"primaryKey"`
	Email        string
	CNPJ         string
	SocialReason string
	User         string
	Password     string
	Agency       Agency `gorm:"foreignKey:AgencyID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
