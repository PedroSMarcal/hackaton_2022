package models

import (
	"time"

	"gorm.io/gorm"
)

type Entrepeneur struct {
	ID          uint `gorm:"primaryKey"`
	AccounterID uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
