package models

import (
	"time"

	"gorm.io/gorm"
)

type Agency struct {
	ID               uint           `gorm:"primaryKey"`
	CNPJ             string         `json:"cnpj"`
	RazaoSocial      string         `json:"razaosocial"`
	Login            string         `json:"login"`
	Senha            string         `json:"senha"`
	RepresentativeID uint           `json:"representative_id"`
	Representative   Representative `gorm:"foreignKey:RepresentativeID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type AgencyRepository interface {
	GetAgency(agency Agency) Agency
}
