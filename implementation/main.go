package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Angry-Potato/go-pay-me/implementation/db"
	"github.com/Angry-Potato/go-pay-me/implementation/schema"
	"github.com/Angry-Potato/go-pay-me/implementation/web"
	"github.com/jinzhu/gorm"
)

func main() {
	DB, err := initDB(&schema.Payment{}, &schema.PaymentAttributes{})

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
	DB = DB.AutoMigrate(models...)
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("internal_payment_id", "payments(id)", "CASCADE", "CASCADE")
	DB.LogMode(true)
	return DB, nil
}
