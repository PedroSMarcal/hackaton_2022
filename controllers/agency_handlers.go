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
	agencyHandler struct{}

	IAgencyHandler interface {
		hello(w http.ResponseWriter, r *http.Request)
	}
)

func NewAgencyHandler() *agencyHandler {
	return &agencyHandler{}
}

func (a *agencyHandler) sendBankReconciliation(w http.ResponseWriter, r *http.Request) {
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

		result, err := service.ConvertTransactionsToBankRelation()
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

func (a *agencyHandler) postAgency(w http.ResponseWriter, r *http.Request) {
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

		agency := models.Agency{}

		value, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid user")
			return
		}

		json.Unmarshal(value, &agency)

		err = repository.PostAgency(&agency)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "falha ao criar usuario")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&agency)

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

func (a *agencyHandler) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

func (a *agencyHandler) postRepresentative(w http.ResponseWriter, r *http.Request) {
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

		representative := models.Representative{}

		value, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid user")
			return
		}

		json.Unmarshal(value, &representative)

		err = repository.PostRepresentative(&representative)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "falha ao criar usuario")
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
