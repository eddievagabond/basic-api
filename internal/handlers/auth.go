package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eddievagabond/internal/models"
	"github.com/eddievagabond/internal/util"

	"github.com/gorilla/mux"
)

type authHandler struct {
	authRepo models.AuthRepository
}

func NewAuthHandler(authRepo models.AuthRepository) *authHandler {
	return &authHandler{
		authRepo,
	}
}

func RegisterAuthHandler(authRepo models.AuthRepository, r *mux.Router) {
	h := NewAuthHandler(authRepo)
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", h.signup).Methods("POST")
	authRouter.HandleFunc("/login", h.login).Methods("POST")
	authRouter.HandleFunc("/refresh", h.refresh).Methods("POST")
}

func (h *authHandler) signup(w http.ResponseWriter, r *http.Request) {
	var u models.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPass, err := util.HashPassword(u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.authRepo.Create(r.Context(), &u, hashedPass)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO return the token and refresh token instead of user
	w.Write(result)
}

func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
	var u models.LoginUserParams
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.authRepo.GetByEmail(r.Context(), u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = util.CheckPasswordHash(u.Password, user.HashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	result, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO return the token and refresh token instead of user
	w.Write(result)
}

func (h *authHandler) refresh(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
