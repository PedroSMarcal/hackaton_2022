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
	login := NewLoginHandler()
	entrepeneur := NewEntrepeneurHandler()

	mux.HandleFunc("/representative", entrepeneur.representativeGet)
	mux.HandleFunc("/representative/create", agency.postRepresentative)

	mux.HandleFunc("/agency/bankreconciliation", agency.sendBankReconciliation)
	mux.HandleFunc("/agency/create", agency.postAgency)
	mux.HandleFunc("/agency/", agency.postAgency)

	mux.HandleFunc("/entrepeneur/cashflow", entrepeneur.cashFlow)
	mux.HandleFunc("/entrepeneur/create", entrepeneur.cashFlow)
	mux.HandleFunc("/entrepeneur/", entrepeneur.get)

	mux.HandleFunc("/login", login.LoginAgencyHandler)

	mux.HandleFunc("/health", agency.hello)
}

func Run() {
	mux := start()

	setRoutes(mux)

	http.ListenAndServe(":"+configs.EnvVariable.Port, mux)
}
