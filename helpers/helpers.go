package helpers

import (
	"fmt"
	"strings"

	"github.com/PedroSMarcal/hackaton2022/common"
)

func ValidateAuthorizationToken(str string) bool {
	if str != common.Token {
		return false
	}

	return true
}

func ValidateEmptyString(str string) bool {
	if str != "" {
		return true
	}
	return false
}

func GetFormatedDSN(dsn string) string {
	credentials := getCredentialsFromDSN(dsn)

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=America/Sao_Paulo",
		credentials.Host,
		credentials.User,
		credentials.Password,
		credentials.Name,
		credentials.Port,
	)

}

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func getCredentialsFromDSN(dsn string) *Config {
	pureDSN := strings.Trim(dsn, "postgres://")
	credentials := strings.Split(pureDSN, "@")

	authCredentials := credentials[0]
	configCredentials := credentials[len(credentials)-1]

	user, password, _ := strings.Cut(authCredentials, `:`)
	hostConfig, name, _ := strings.Cut(configCredentials, "/")
	host, port, _ := strings.Cut(hostConfig, ":")

	return &Config{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name,
	}
}
