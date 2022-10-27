package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVar struct {
	Port string
	Url  string
}

var EnvVariable *EnvironmentVar

func GetEnvironmentVariables() *EnvironmentVar {
	godotenv.Load(".env")

	EnvVariable = &EnvironmentVar{
		Url:  os.Getenv("DATABASE_URL"),
		Port: os.Getenv("PORT"),
	}
	return EnvVariable
}
