package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eddievagabond/internal/handlers"
	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	router := mux.NewRouter()
	handlers.RegisterHealthHandler(router)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/health", nil)

	router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
