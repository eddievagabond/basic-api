package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/eddievagabond/internal/handlers"
	"github.com/eddievagabond/internal/middleware"
	"github.com/eddievagabond/internal/storage"

	"github.com/gorilla/mux"
)

var defaultStipTimeout = time.Second * 10

type Server struct {
	addr    string
	storage *storage.Storage
}

func NewApiServer(addr string, storage *storage.Storage) (*Server, error) {
	if addr == "" {
		return nil, errors.New("address is empty")
	}

	return &Server{
		addr:    addr,
		storage: storage,
	}, nil
}

func (s *Server) Start(stop chan struct{}) error {
	svr := &http.Server{
		Addr:    s.addr,
		Handler: s.router(),
	}

	go func() {
		log.Printf("Starting server on %s", s.addr)
		if err := svr.ListenAndServe(); err != nil {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), defaultStipTimeout)
	defer cancel()
	log.Printf("Timeout on %s", s.addr)

	return svr.Shutdown(ctx)
}

func (s *Server) router() http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.ResponseHeaderMiddleware)

	handlers.RegisterProductsHandler(s.storage.ProductRepository, router)
	handlers.RegisterHealthHandler(router)

	return router
}
