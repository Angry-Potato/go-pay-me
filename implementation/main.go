package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Angry-Potato/go-pay-me/implementation/db"
	"github.com/Angry-Potato/go-pay-me/implementation/schema"
	"github.com/Angry-Potato/go-pay-me/implementation/web"
)

func main() {
	DB, err := db.Initialise(&schema.Payment{}, &schema.PaymentAttributes{}, &schema.Party{}, &schema.CurrencyExchange{})
	if err != nil {
		log.Fatalf("Error initialising database: %s", err.Error())
	}

	//this is bad, where should I do this?
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("internal_payment_id", "payments(id)", "CASCADE", "CASCADE")
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("beneficiary_party_id", "parties(id)", "SET NULL", "CASCADE")
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("debtor_party_id", "parties(id)", "SET NULL", "CASCADE")
	DB.Model(&schema.PaymentAttributes{}).AddForeignKey("sponsor_party_id", "parties(id)", "SET NULL", "CASCADE")
	DB.Model(&schema.CurrencyExchange{}).AddForeignKey("payment_attributes_id", "payment_attributes(id)", "CASCADE", "CASCADE")

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
