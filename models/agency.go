package models

import (
	"time"

	"gorm.io/gorm"
)

type Agency struct {
	ID            uint `gorm:"primaryKey"`
	CNPJ          string
	RazaoSocial   string
	EntrepeneurID uint
	Entrepeneur   []Entrepeneur `gorm:"foreignKey:EntrepeneurID"`
	Credencials   Credencials
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type AgencyRepository interface {
	GetAgency(agency Agency) Agency
}
