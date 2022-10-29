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

func (l *handlerLogin) LoginHandler(w http.ResponseWriter, r *http.Request) {
	credentials := models.Login{}

	values, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(values, &credentials)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "invalid user")
	}

	if service.Login(credentials.User, credentials.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "invalid user")
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "bbad700841c8f8568dc3d46d737bb24a7752aa348b6fd26aaee2d178f177a532")
}
