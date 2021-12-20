package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eddievagabond/internal/server/models"
	"github.com/gorilla/mux"
)

type handler struct {
	productRepo models.ProductRepository
}

func New(productRepo models.ProductRepository) *handler {
	return &handler{
		productRepo,
	}
}

func RegisterProductsHandler(productRepo models.ProductRepository, r *mux.Router) {
	h := New(productRepo)
	productsRouter := r.PathPrefix("/products").Subrouter()

	productsRouter.HandleFunc("", h.get).Methods("GET")
	productsRouter.HandleFunc("/{id}", h.getById).Methods("GET")
	productsRouter.HandleFunc("", h.create).Methods("POST")
	productsRouter.HandleFunc("", h.update).Methods("PUT")
	productsRouter.HandleFunc("/{id}", h.delete).Methods("DELETE")
}

func (h *handler) getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := h.productRepo.GetById(r.Context(), id)

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

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	count, err := strconv.Atoi(query.Get("count"))
	if err != nil {
		count = 10
	}
	start, err := strconv.Atoi(query.Get("start"))
	if err != nil {
		start = 0
	}

	data, err := h.productRepo.Get(r.Context(), start, count)

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

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	p := &models.Product{}

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
	}

	data, err := h.productRepo.Create(context.Background(), p)

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

func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	p := &models.Product{}

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
	}

	data, err := h.productRepo.Update(context.Background(), p)

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

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.productRepo.Delete(context.Background(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
