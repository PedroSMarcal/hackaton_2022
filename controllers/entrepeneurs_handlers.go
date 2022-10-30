package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/PedroSMarcal/hackaton2022/helpers"
	"github.com/PedroSMarcal/hackaton2022/service"
)

type (
	entrepenneurHandler struct{}

	IEntrepenneurHandler interface {
		hello(w http.ResponseWriter, r *http.Request)
		cashFlow(w http.ResponseWriter, r *http.Request)
	}
)

func NewEntrepeneurHandler() *entrepenneurHandler {
	return &entrepenneurHandler{}
}

func (a *entrepenneurHandler) cashFlow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	switch r.Method {
	case http.MethodGet:
		values := r.URL.Query()
		token := values.Get("Authorization")

		valid := helpers.ValidateAuthorizationToken(token)
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid user")
			return
		}

		result, err := service.ConvertTransactionToCashFlowReturn()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid user")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&result)

	case http.MethodOptions:
		w.WriteHeader(http.StatusAccepted)
		io.WriteString(w, "Bem Vindo")
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not Allowed")
		return
	}
}
