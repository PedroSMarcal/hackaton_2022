package models

import (
	"time"

	"gorm.io/gorm"
)

type Agency struct {
	ID          uint `gorm:"primaryKey"`
	CNPJ        uint
	RazaoSocial uint
	Accounts    []Accounts `gorm:"foreignKey:AccountsID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type AgencyRepository interface {
	GetAgency(agency Agency) Agency
}
