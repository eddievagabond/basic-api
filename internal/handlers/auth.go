package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eddievagabond/internal/models"
	"github.com/eddievagabond/internal/services"

	"github.com/gorilla/mux"
)

type authHandler struct {
	authService services.AuthenticationService
}

func NewAuthHandler(authService services.AuthenticationService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}

func RegisterAuthHandler(authService services.AuthenticationService, r *mux.Router) {
	h := NewAuthHandler(authService)
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", h.register).Methods("POST")
	authRouter.HandleFunc("/authenticate", h.authenticate).Methods("POST")
	authRouter.HandleFunc("/refresh", h.refresh).Methods("POST")
}

func (h *authHandler) register(w http.ResponseWriter, r *http.Request) {
	var u models.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.authService.Register(r.Context(), &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, err := h.authService.GenerateAccessToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshToken, err := h.authService.GenerateRefreshToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ar := models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	result, err := json.Marshal(ar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (h *authHandler) authenticate(w http.ResponseWriter, r *http.Request) {
	var u models.LoginUserParams
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.authService.Authenticate(r.Context(), u.Email, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	accessToken, err := h.authService.GenerateAccessToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshToken, err := h.authService.GenerateRefreshToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ar := models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	result, err := json.Marshal(ar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (h *authHandler) refresh(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
