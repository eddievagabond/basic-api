package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeProductHandlers(r *mux.Router) {
	productsSub := r.PathPrefix("/products").Subrouter()
	productsSub.HandleFunc("", getProducts).Methods("GET")
	productsSub.HandleFunc("/{id:[0-9]+}", getProduct).Methods("GET")
	productsSub.HandleFunc("", createProduct).Methods("POST")
	productsSub.HandleFunc("/{id:[0-9]+}", updateProduct).Methods("PUT")
	productsSub.HandleFunc("/{id:[0-9]+}", deleteProduct).Methods("DELETE")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
