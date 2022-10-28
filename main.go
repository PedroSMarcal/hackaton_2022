package main

import (
	"github.com/PedroSMarcal/hackaton2022/configs"
	"github.com/PedroSMarcal/hackaton2022/controllers"
)

func main() {
	configs.GetEnvironmentVariables()

	controllers.Run()
}
