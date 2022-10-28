package controllers

import "net/http"

type (
	agencyHandler struct{}

	IAgencyHandler interface {
		hello(w http.ResponseWriter, r *http.Request)
	}
)

func NewAgencyHandler() *agencyHandler {
	return &agencyHandler{}
}

func (a *agencyHandler) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
