package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func ResponseHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

var UnauthenticatedRoutes = [3]string{
	"/auth/authenticate",
	"/auth/register",
	"/auth/refresh",
}

func (h *AuthHandler) ValidateAccessTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range UnauthenticatedRoutes {
			if route == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}

		token, err := extractToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userId, err := h.authService.ValidateAccessToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey{}, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		return "", fmt.Errorf("authorization header format is invalid")
	}
	return authHeaderContent[1], nil
}
