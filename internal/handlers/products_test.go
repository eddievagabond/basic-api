package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eddievagabond/internal/handlers"
	"github.com/eddievagabond/internal/models"
	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
)

type mockProductsRepository struct{}

func (m mockProductsRepository) Get(ctx context.Context, start, count int) ([]*models.Product, error) {
	products := []*models.Product{
		{ID: "1", Name: "Product 1", Price: 1.0},
		{ID: "2", Name: "Product 2", Price: 2.0},
	}

	return products, nil
}

func (m mockProductsRepository) GetById(ctx context.Context, id string) (*models.Product, error) {
	p := &models.Product{ID: id}

	return p, nil
}

func (r mockProductsRepository) Create(ctx context.Context, p *models.Product) (*models.Product, error) {

	return p, nil
}

func (m mockProductsRepository) Update(ctx context.Context, p *models.Product) (*models.Product, error) {
	return p, nil
}

func (m mockProductsRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func TestProductsHandlerGet(t *testing.T) {
	router := mux.NewRouter()
	repo := mockProductsRepository{}
	handlers.RegisterProductsHandler(repo, router)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/products", nil)

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	var result []*models.Product
	err := json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
}

func TestProductsHandlerGetById(t *testing.T) {
	router := mux.NewRouter()
	repo := mockProductsRepository{}
	handlers.RegisterProductsHandler(repo, router)

	p := models.Product{ID: "1"}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", fmt.Sprintf("/products/%s", p.ID), nil)

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	var result models.Product
	err := json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, p.ID, result.ID)
}

func TestProductsHandlerCreate(t *testing.T) {
	router := mux.NewRouter()
	repo := mockProductsRepository{}
	handlers.RegisterProductsHandler(repo, router)

	p := models.Product{ID: "1", Name: "Product 1", Price: 1.0}
	w := httptest.NewRecorder()
	pBytes, _ := json.Marshal(&p)

	r := httptest.NewRequest("POST", "/products", bytes.NewBuffer(pBytes))

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	var result models.Product
	err := json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, p.ID, result.ID)
}

func TestProductsHandlerUpdate(t *testing.T) {
	router := mux.NewRouter()
	repo := mockProductsRepository{}
	handlers.RegisterProductsHandler(repo, router)

	p := models.Product{ID: "1", Name: "Product 1", Price: 1.0}
	w := httptest.NewRecorder()
	pBytes, _ := json.Marshal(&p)

	r := httptest.NewRequest("PUT", "/products", bytes.NewBuffer(pBytes))

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	var result models.Product
	err := json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, p.ID, result.ID)
}

func TestProductsDelete(t *testing.T) {
	router := mux.NewRouter()
	repo := mockProductsRepository{}
	handlers.RegisterProductsHandler(repo, router)

	p := models.Product{ID: "1"}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", fmt.Sprintf("/products/%s", p.ID), nil)

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
