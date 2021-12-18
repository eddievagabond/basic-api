package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eddievagabond/internal/storage"
	"github.com/gorilla/mux"
)

type handler struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *handler {
	return &handler{
		storage: storage,
	}
}

func RegisterProductsHandler(s *storage.Storage, r *mux.Router) {
	h := New(s)
	productsRouter := r.PathPrefix("/products").Subrouter()

	productsRouter.HandleFunc("", h.listProducts).Methods("GET")
	productsRouter.HandleFunc("/{id}", h.getProduct).Methods("GET")
	productsRouter.HandleFunc("", h.createProduct).Methods("POST")
	productsRouter.HandleFunc("", h.updateProduct).Methods("PUT")
	productsRouter.HandleFunc("/{id}", h.deleteProduct).Methods("DELETE")
}

func (h *handler) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := h.storage.GetProduct(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(result)
}

func (h *handler) listProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	count, err := strconv.Atoi(query.Get("count"))
	if err != nil {
		count = 10
	}
	start, err := strconv.Atoi(query.Get("start"))
	if err != nil {
		start = 0
	}

	data, err := h.storage.ListProducts(r.Context(), start, count)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(result)
}

func (s *handler) createProduct(w http.ResponseWriter, r *http.Request) {
	pr := storage.CreateProductRequest{}

	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
	}

	data, err := s.storage.CreateProduct(context.Background(), pr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(result)
}

func (h *handler) updateProduct(w http.ResponseWriter, r *http.Request) {
	p := storage.Product{}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
	}

	data, err := h.storage.UpdateProduct(context.Background(), p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(result)
}

func (h *handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.storage.DeleteProduct(context.Background(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
