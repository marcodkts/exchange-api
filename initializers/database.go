package initializers

import (
	"log"
	"os"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var host = os.Getenv("DB_HOST")
	var user = os.Getenv("DB_USER")
	var pass = os.Getenv("DB_PASSWORD")
	var name = os.Getenv("DB_NAME")
	var port = os.Getenv("DB_PORT")
	var err error

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable`, host, user, pass, name, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
}