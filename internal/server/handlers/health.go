package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHealthHandler(r *mux.Router) {
	r.HandleFunc("/health", healthHandler)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
