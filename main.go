package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ElladanTasartir/go-bank-api/config"
	"github.com/ElladanTasartir/go-bank-api/db"
	"github.com/ElladanTasartir/go-bank-api/routes"
)

func main() {
	envs := config.LoadEnv()
	port := envs["PORT"]
	host := envs["POSTGRES_HOST"]
	user := envs["POSTGRES_USER"]
	password := envs["POSTGRES_PASSWORD"]
	database := envs["POSTGRES_DB"]
	dbPort, err := strconv.Atoi(envs["POSTGRES_PORT"])
	if err != nil {
		log.Fatal("Port is not a number")
	}

	connection := db.InitDatabase(host, user, password, database, dbPort)
	log.Printf("Database connected")
	log.Printf("Listening on port %v \n", port)
	address := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(address, routes.NewRouter(connection)))
}
