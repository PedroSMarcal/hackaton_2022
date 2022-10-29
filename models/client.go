package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID        uint       `gorm:"primaryKey"`
	Accounts  []Accounts `gorm:"foreignKey:AccountsID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
