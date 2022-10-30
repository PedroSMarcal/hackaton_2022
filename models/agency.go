package models

import (
	"time"

	"gorm.io/gorm"
)

type Agency struct {
	ID            uint `gorm:"primaryKey"`
	CNPJ          uint
	RazaoSocial   uint
	EntrepeneurID uint
	Entrepeneur   []Entrepeneur `gorm:"foreignKey:EntrepeneurID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type AgencyRepository interface {
	GetAgency(agency Agency) Agency
}
