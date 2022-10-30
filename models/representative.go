package models

import (
	"time"

	"gorm.io/gorm"
)

type Representative struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `json:"email"`
	CNPJ         string `json:"cnpj"`
	SocialReason string `json:"socialreason"`
	Login        string `json:"login"`
	Senha        string `json:"senha"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
