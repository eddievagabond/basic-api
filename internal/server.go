package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/eddievagabond/internal/handlers"
	"github.com/eddievagabond/internal/middleware"
	"github.com/eddievagabond/internal/services"
	"github.com/eddievagabond/internal/storage"
	"github.com/eddievagabond/internal/util"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var defaultStipTimeout = time.Second * 10

type Server struct {
	config  *util.Configuration
	storage *storage.Storage
}

func NewApiServer(config *util.Configuration, storage *storage.Storage) (*Server, error) {
	if config.ServerAddress == "" {
		return nil, errors.New("server address is empty in config")
	}

	return &Server{
		config:  config,
		storage: storage,
	}, nil
}

func (s *Server) Start(stop chan struct{}) error {
	svr := &http.Server{
		Addr:    s.config.ServerAddress,
		Handler: s.router(),
	}

	go func() {
		log.Printf("Starting server on %s", s.config.ServerAddress)
		if err := svr.ListenAndServe(); err != nil {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), defaultStipTimeout)
	defer cancel()
	log.Printf("Timeout on %s", s.config.ServerAddress)

	return svr.Shutdown(ctx)
}

func (s *Server) router() http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.ResponseHeaderMiddleware)

	authService := services.NewAuthService(s.config, s.storage.UserRepository)

	handlers.RegisterAuthHandler(authService, router)
	handlers.RegisterUserHandler(s.storage.UserRepository, router)
	handlers.RegisterProductsHandler(s.storage.ProductRepository, router)
	handlers.RegisterTransferHandler(s.storage.TransferRepository, router)
	handlers.RegisterAccountHandler(s.storage.AccountRepository, router)
	handlers.RegisterHealthHandler(router)

	// TODO: Get from env config
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		Debug:            false,
	})

	handler := c.Handler(router)

	return handler
}
