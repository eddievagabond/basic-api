package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eddievagabond/internal/models"
	"github.com/gorilla/mux"
)

type transferHandler struct {
	transferRepo models.TransferRepository
}

func NewTransferHandler(transferRepo models.TransferRepository) *transferHandler {
	return &transferHandler{
		transferRepo,
	}
}

func RegisterTransferHandler(transferRepo models.TransferRepository, r *mux.Router) {
	h := NewTransferHandler(transferRepo)
	transferRouter := r.PathPrefix("/transfer").Subrouter()

	transferRouter.HandleFunc("", h.get).Methods("GET")
	transferRouter.HandleFunc("/{id}", h.getById).Methods("GET")
	transferRouter.HandleFunc("", h.create).Methods("POST")
	// transferRouter.HandleFunc("", h.update).Methods("PUT")
	// transferRouter.HandleFunc("/{id}", h.delete).Methods("DELETE")
}

func (h *transferHandler) get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	count, err := strconv.Atoi(query.Get("count"))
	if err != nil {
		count = 10
	}
	start, err := strconv.Atoi(query.Get("start"))
	if err != nil {
		start = 0
	}

	data, err := h.transferRepo.Get(r.Context(), start, count)

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

func (h *transferHandler) getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := h.transferRepo.GetById(r.Context(), id)

	if data == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

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

func (h *transferHandler) create(w http.ResponseWriter, r *http.Request) {
	t := &models.Transfer{}

	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := h.transferRepo.TransferTx(r.Context(), t)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}
