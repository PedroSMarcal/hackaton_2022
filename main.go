package main

import (
	"github.com/PedroSMarcal/hackaton2022/configs"
	"github.com/PedroSMarcal/hackaton2022/repository/database"
)

func main() {
	configs.GetEnvironmentVariables()

	database.StartDb()
	// controllers.Run()
}
