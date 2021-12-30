package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewAuthHandler() {

}

func RegisterAuthHandler(r *mux.Router) {
	NewAuthHandler()
	r.HandleFunc("/auth/{provider}/callback", handleProviderCallback)
	r.HandleFunc("/auth/{provider}", handleProviderBegin)
}

func handleProviderCallback(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleProviderBegin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
