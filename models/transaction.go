package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            uint        `gorm:"primaryKey"`
	Total         int         `json:"total"`
	TotalPages    int         `json:"totalPages"`
	Page          int         `json:"page"`
	EntrepeneurID uint        `json:"entrepeneur_id"`
	Entrepeneur   Entrepeneur `gorm:"foreignKey:EntrepeneurID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
