package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eddievagabond/internal/models"
	"github.com/eddievagabond/internal/services"

	"github.com/gorilla/mux"
)

// UserIDKey is used as a key for storing the UserID in context at middleware
type UserIDKey struct{}

type AuthHandler struct {
	authService services.AuthenticationService
}

func NewAuthHandler(authService services.AuthenticationService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func RegisterAuthHandler(h *AuthHandler, r *mux.Router) {
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", h.register).Methods("POST")
	authRouter.HandleFunc("/authenticate", h.authenticate).Methods("POST")
	authRouter.HandleFunc("/refresh", h.refresh).Methods("POST")
}

func (h *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
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

func (h *AuthHandler) authenticate(w http.ResponseWriter, r *http.Request) {
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

func (h *AuthHandler) refresh(w http.ResponseWriter, r *http.Request) {
	var rrq models.RefreshRequestParams
	err := json.NewDecoder(r.Body).Decode(&rrq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := h.authService.ValidateRefreshToken(rrq.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	accessToken, err := h.authService.GenerateAccessToken(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshToken, err := h.authService.GenerateRefreshToken(userId)
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
