package database

import (
	"log"

	"github.com/PedroSMarcal/hackaton2022/configs"
	"github.com/PedroSMarcal/hackaton2022/helpers"
	"github.com/PedroSMarcal/hackaton2022/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb() {
	dsn := helpers.GetFormatedDSN(configs.EnvVariable.Url)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db = database

	db.AutoMigrate(&models.Agency{})
}

func OpenConnection() *gorm.DB {
	StartDb()

	return db
}
