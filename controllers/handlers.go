package controllers

import (
	"net/http"

	"github.com/PedroSMarcal/hackaton2022/configs"
)

func start() *http.ServeMux {
	// mux := http.NewServeMux()
	mux := http.ServeMux{}

	return &mux
}

func setRoutes(mux *http.ServeMux) {
	agency := NewAgencyHandler()

	mux.HandleFunc("/health", agency.hello)
}

func Run() {
	mux := start()

	setRoutes(mux)

	http.ListenAndServe(":"+configs.EnvVariable.Port, mux)
}
