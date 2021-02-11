package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	bank := NewBank()
	router.Handle("/banks", bank.router)
	return router
}
