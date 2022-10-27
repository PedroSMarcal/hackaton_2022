package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVar struct {
	DbUser       string
	DbPassword   string
	DbHost       string
	DbPort       string
	DbName       string
	Port         string
	JwtSignature string
}

var EnvVariable *EnvironmentVar

func GetEnvironmentVariables() *EnvironmentVar {
	godotenv.Load(".env")

	EnvVariable = &EnvironmentVar{
		DbUser:       os.Getenv("DB_USER"),
		DbPassword:   os.Getenv("DB_PASSWORD"),
		DbHost:       os.Getenv("DB_HOST"),
		DbPort:       os.Getenv("DB_PORT"),
		DbName:       os.Getenv("DB_NAME"),
		Port:         os.Getenv("PORT"),
		JwtSignature: os.Getenv("JWT_SIGNATURE"),
	}
	return EnvVariable
}
