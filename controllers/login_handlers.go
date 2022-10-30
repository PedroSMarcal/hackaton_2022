package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/PedroSMarcal/hackaton2022/models"
	"github.com/PedroSMarcal/hackaton2022/service"
)

type (
	handlerLogin struct {
	}

	ILoginHandler interface {
		LoginHandler(w http.ResponseWriter, r *http.Request)
	}
)

func NewLoginHandler() *handlerLogin {
	return &handlerLogin{}
}

func (l *handlerLogin) LoginAgencyHandler(w http.ResponseWriter, r *http.Request) {

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusAccepted)
		io.WriteString(w, "Ok")
	}

	credentials := models.Login{}

	values, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(values, &credentials)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "invalid user")
		return
	}

	valid := service.Login(credentials.User, credentials.Password)

	if valid == "" {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "invalid user")
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, valid)
}
