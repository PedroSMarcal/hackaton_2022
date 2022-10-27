package database

import (
	"fmt"
	"log"

	"github.com/PedroSMarcal/hackaton2022/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb(
	DbUser string,
	DbPassword string,
	DbHost string,
	DbPort string,
	DbName string,
) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName,
	)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db = database

}

func OpenConnection() *gorm.DB {
	StartDb(
		configs.EnvVariable.DbUser,
		configs.EnvVariable.DbPassword,
		configs.EnvVariable.DbHost,
		configs.EnvVariable.DbPort,
		configs.EnvVariable.DbName,
	)

	return db
}
