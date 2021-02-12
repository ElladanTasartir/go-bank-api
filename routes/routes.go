package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(connection *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	bank := NewBank(connection)
	router.Handle("/banks", bank.router)
	return router
}
