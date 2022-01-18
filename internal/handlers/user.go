package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eddievagabond/internal/models"
	"github.com/gorilla/mux"
)

type userHandler struct {
	userRepo models.UserRepository
}

func NewUserHandler(userRepo models.UserRepository) *userHandler {
	return &userHandler{
		userRepo,
	}
}

func RegisterUserHandler(userRepo models.UserRepository, r *mux.Router) {
	h := NewUserHandler(userRepo)
	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/profile/{id}", h.getProfile).Methods("GET")
}

func (h *userHandler) getProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := h.userRepo.GetById(r.Context(), id)
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
