package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type BankRouter struct {
	router  *mux.Router
	banks   []Bank
	idCount int
}

type Bank struct {
	ID   int
	Name string `json:"name"`
	Code int    `json:"code"`
}

func (b *BankRouter) registerRoutes() {
	b.router.HandleFunc("/banks", b.getBanks).Methods("GET")
	b.router.HandleFunc("/banks", b.createBank).Methods("POST")
}

func (b *BankRouter) getBanks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(b.banks)
}

func (b *BankRouter) createBank(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bank := &Bank{
		ID: b.idCount + 1,
	}

	b.idCount++
	json.Unmarshal(body, bank)
	b.banks = append(b.banks, *bank)

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(*bank)
}

func NewBank() *BankRouter {
	bank := BankRouter{
		router: mux.NewRouter(),
		banks:  []Bank{},
	}

	bank.registerRoutes()

	return &bank
}
