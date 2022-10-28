package controllers

import (
	"net/http"

	"github.com/PedroSMarcal/hackaton2022/configs"
)

func Start() {
	mux := http.NewServeMux()

	http.ListenAndServe(configs.EnvVariable.Port, mux)
}
