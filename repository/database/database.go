package database

import (
	"log"

	"github.com/PedroSMarcal/hackaton2022/configs"
	"github.com/PedroSMarcal/hackaton2022/helpers"
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

	// db.AutoMigrate(&	models.Transaction{}, &models.Representative{}, &models.Entrepeneur{}, &models.Association{}, &models.Agency{})
}

func OpenConnection() *gorm.DB {
	if db == nil {
		StartDb()
	}

	return db
}
