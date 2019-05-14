package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Angry-Potato/go-pay-me/implementation/payments"

	"github.com/Angry-Potato/go-pay-me/implementation/db"
	"github.com/Angry-Potato/go-pay-me/implementation/web"
	"github.com/jinzhu/gorm"
)

func main() {
	DB, err := initDB(&payments.Payment{})

	if err != nil {
		log.Fatalf("Error initialising database: %s", err.Error())
	}

	serverPort := port(os.Getenv("PORT"), 8080)
	log.Fatal(web.StartServer(serverPort, DB))
}

func port(envPort string, defaultPort int) int {
	if envPort == "" {
		return defaultPort
	}

	parsedPort, _ := strconv.Atoi(envPort)
	return parsedPort
}

func initDB(models ...interface{}) (*gorm.DB, error) {
	DB, err := db.Connect(
		os.Getenv("DATABASE_URL"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		return nil, err
	}

	return DB.AutoMigrate(models...), nil
}
