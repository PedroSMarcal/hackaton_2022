package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/PedroSMarcal/hackaton2022/helpers"
	"github.com/PedroSMarcal/hackaton2022/models"
	"github.com/PedroSMarcal/hackaton2022/repository"
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
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
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

func (a *entrepenneurHandler) get(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
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

		entrepeneur := &[]models.Entrepeneur{}

		err := repository.GetAllEntrepeneur(entrepeneur)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid user")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&entrepeneur)

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

func (a *entrepenneurHandler) representativeGet(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
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

		representative := &[]models.Representative{}

		err := repository.GetRepresentative(representative)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid user")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&representative)

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

func (a *entrepenneurHandler) entrepeteursGet(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
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

		entrepeneur := &[]models.Entrepeneur{}

		err := repository.GetAllEntrepeneur(entrepeneur)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid user")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&entrepeneur)

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

func (a *agencyHandler) postEntrepeteurs(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	switch r.Method {
	case http.MethodPost:
		values := r.URL.Query()
		token := values.Get("Authorization")

		valid := helpers.ValidateAuthorizationToken(token)
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid user")
			return
		}

		entrepeneur := models.Entrepeneur{}

		value, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid user")
			return
		}

		json.Unmarshal(value, &entrepeneur)

		err = repository.PostEntrepeteurs(&entrepeneur)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "falha ao criar usuario")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&entrepeneur)

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
