package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ElladanTasartir/go-bank-api/config"
	"github.com/ElladanTasartir/go-bank-api/routes"
)

func main() {
	config.LoadEnv()
	port := os.Getenv("PORT")
	log.Printf("Listening on port %v \n", port)
	address := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(address, routes.NewRouter()))
}
