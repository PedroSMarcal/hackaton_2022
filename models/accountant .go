package models

import (
	"time"

	"gorm.io/gorm"
)

type Accounts struct {
	ID          uint `gorm:"primaryKey"`
	AgencyID    uint
	Name        string
	Email       string
	Entrepeneur []Entrepeneur `gorm:"foreignKey:EntrepeneurID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
