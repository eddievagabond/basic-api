package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eddievagabond/internal/models"
	"github.com/gorilla/mux"
)

type accountHandler struct {
	accountRepo models.AccountRepository
}

func NewAccountHandler(accountRepo models.AccountRepository) *accountHandler {
	return &accountHandler{
		accountRepo,
	}
}

func RegisterAccountHandler(accountRepo models.AccountRepository, r *mux.Router) {
	h := NewAccountHandler(accountRepo)
	accountRouter := r.PathPrefix("/account").Subrouter()

	accountRouter.HandleFunc("", h.get).Methods("GET")
	accountRouter.HandleFunc("/{id}", h.getById).Methods("GET")
	accountRouter.HandleFunc("", h.create).Methods("POST")
	accountRouter.HandleFunc("", h.update).Methods("PUT")
	accountRouter.HandleFunc("/{id}", h.delete).Methods("DELETE")
}

func (h *accountHandler) get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	count, err := strconv.Atoi(query.Get("count"))
	if err != nil {
		count = 10
	}
	start, err := strconv.Atoi(query.Get("start"))
	if err != nil {
		start = 0
	}

	data, err := h.accountRepo.Get(r.Context(), start, count)

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

func (h *accountHandler) getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := h.accountRepo.GetById(r.Context(), id)

	if data == nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

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

func (h *accountHandler) create(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := h.accountRepo.Create(r.Context(), &account)

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

func (h *accountHandler) update(w http.ResponseWriter, r *http.Request) {
	a := &models.Account{}
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}

	data, err := h.accountRepo.Update(r.Context(), a)

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

func (h *accountHandler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.accountRepo.Delete(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
