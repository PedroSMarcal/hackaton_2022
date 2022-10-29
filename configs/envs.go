package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVar struct {
	Port         string
	Url          string
	ClientID     string
	ClientSecret string
}

var EnvVariable *EnvironmentVar

func GetEnvironmentVariables() *EnvironmentVar {
	godotenv.Load(".env")

	EnvVariable = &EnvironmentVar{
		Url:          os.Getenv("DATABASE_URL"),
		Port:         os.Getenv("PORT"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
	return EnvVariable
}
