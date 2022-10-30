package repository

import (
	"github.com/PedroSMarcal/hackaton2022/models"
	"github.com/PedroSMarcal/hackaton2022/repository/database"
)

func PostAgency(agency *models.Agency) error {
	db := database.OpenConnection()

	return db.Create(&agency).Error
}

func GetAgency(agency *models.Agency) error {
	db := database.OpenConnection()

	return db.Find(&agency).Error
}

func PostRepresentative(representative *models.Representative) error {
	db := database.OpenConnection()

	return db.Create(&representative).Error
}

func GetRepresentative(representative *models.Representative) error {
	db := database.OpenConnection()

	return db.Find(&representative).Error
}

func PostAssociation(association *models.Association) error {
	db := database.OpenConnection()

	return db.Create(&association).Error
}

func GetByRepresentative(association *models.Association) error {
	db := database.OpenConnection()

	return db.Find(&association).Error
}
