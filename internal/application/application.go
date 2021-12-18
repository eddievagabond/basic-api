package application

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/eddievagabond/internal/application/config"
	"github.com/eddievagabond/internal/application/handler"
	"github.com/eddievagabond/internal/application/storage"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Application struct {
	router  *mux.Router
	storage *storage.Storage
}

func Initialize() *Application {
	dbConfig := config.GetDBConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Schema, dbConfig.SSLMode)
	s, err := storage.NewStorage(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	r := initializeRouter()
	return &Application{
		router:  r,
		storage: s,
	}
}

func (a *Application) Run() {
	serverConfig := config.GetServerConfig()
	addr := fmt.Sprintf("%s:%s", serverConfig.Address, serverConfig.Port)

	log.Printf("Listening on %s\n", addr)
	if err := http.ListenAndServe(addr, a.router); err != nil {
		log.Fatal(err)
	}
}

func initializeRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", handleHealthCheck).Methods("GET")
	handler.InitializeHandlers(router)

	return router
}

func initializeDB() (*sql.DB, error) {
	dbConfig := config.GetDBConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Schema, dbConfig.SSLMode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
