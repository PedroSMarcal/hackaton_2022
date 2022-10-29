package controllers

import "net/http"

type (
	entrepenneurHandler struct{}

	IEntrepenneurHandler interface {
		hello(w http.ResponseWriter, r *http.Request)
	}
)

func NewEntrepeneurHandler() *agencyHandler {
	return &agencyHandler{}
}

func (a *entrepenneurHandler) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
