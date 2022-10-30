package models

import (
	"time"

	"gorm.io/gorm"
)

type Entrepeneur struct {
	ID             uint   `gorm:"primaryKey"`
	CNPJ           string `json:"cnpj"`
	SocialReason   string `json:"socialreason"`
	Login          string `json:"login"`
	Senha          string `json:"senha"`
	Permission     bool   `json:"permission"`
	AgencyID       uint
	Agency         Agency `gorm:"foreignKey:AgencyID"`
	TimePermission time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
