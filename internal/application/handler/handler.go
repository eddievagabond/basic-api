package handler

import (
	"github.com/gorilla/mux"
)

func InitializeHandlers(r *mux.Router) {
	InitializeProductHandlers(r)
}
