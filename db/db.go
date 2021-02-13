package db

import (
	"fmt"
	"log"

	"github.com/ElladanTasartir/go-bank-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(host string, user string, password string, database string, port int) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to connect to database")
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	db.AutoMigrate(&models.Bank{})
	return db
}
