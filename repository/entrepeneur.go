package repository

import (
	"github.com/PedroSMarcal/hackaton2022/models"
	"github.com/PedroSMarcal/hackaton2022/repository/database"
)

func GetAllEntrepeneur(entrepeneur *[]models.Entrepeneur) error {
	db := database.OpenConnection()

	err := db.Find(&entrepeneur).Error
	if err != nil {
		return err
	}

	return nil
}
