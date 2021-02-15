package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ElladanTasartir/go-bank-api/models"
	"github.com/ElladanTasartir/go-bank-api/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type BankRouter struct {
	router      *mux.Router
	bankservice *services.BankService
}

func (b *BankRouter) registerRoutes() {
	b.router.HandleFunc("/banks", b.getBanks).Methods("GET")
	b.router.HandleFunc("/banks", b.createBank).Methods("POST")
	b.router.HandleFunc("/banks/{id}", b.getBank).Methods("GET")
	b.router.HandleFunc("/banks/{id}", b.deleteBank).Methods("DELETE")
}

func (b *BankRouter) getBanks(w http.ResponseWriter, r *http.Request) {
	banks := &[]models.Bank{}
	err := b.bankservice.GetBanks(banks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(banks)
}

func (b *BankRouter) getBank(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]

	bank := &models.Bank{}
	err := b.bankservice.GetBank(bank, param)
	if err != nil {
		http.Error(w, "Bank with ID %v not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bank)
}

func (b *BankRouter) createBank(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bank := &models.Bank{}
	err = json.Unmarshal(body, bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = b.bankservice.CreateBank(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(*bank)
}

func (b *BankRouter) deleteBank(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]

	bank := &models.Bank{}
	err := b.bankservice.DeleteBanks(bank, param)
	if err != nil {
		http.Error(w, fmt.Sprintf("Bank with ID %v not found", param), http.StatusNotFound)
	}

	w.WriteHeader(204)
	json.NewEncoder(w).Encode(nil)
}

func NewBank(connection *gorm.DB) *BankRouter {
	bank := BankRouter{
		router:      mux.NewRouter(),
		bankservice: services.NewBankService(connection),
	}

	bank.registerRoutes()
	log.Println("Bank routes initialized")

	return &bank
}
