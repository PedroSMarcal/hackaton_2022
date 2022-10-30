package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	Total         int  `json:"total"`
	TransactionID uint
	TotalPages    int `json:"totalPages"`
	Page          int `json:"page"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
