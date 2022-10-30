package models

import (
	"time"

	"gorm.io/gorm"
)

type Entrepeneur struct {
	ID             uint `gorm:"primaryKey"`
	CNPJ           string
	SocialReason   string
	AgencyID       uint
	Transaction    []Transaction `gorm:"foreignKey:AgencyID"`
	Permission     bool
	TimePermission time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
