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
	Credencials  Credencials
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
