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
	DB, err := db.Initialise(&schema.Payment{}, &schema.PaymentAttributes{}, &schema.Party{})

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
