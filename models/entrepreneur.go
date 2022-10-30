package models

import (
	"time"

	"gorm.io/gorm"
)

type Entrepeneur struct {
	ID             uint `gorm:"primaryKey"`
	CNPJ           string
	SocialReason   string
	Permission     bool
	TransactionID  uint
	Transaction    Transaction `gorm:"foreignKey:TransactionID"`
	TimePermission time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
