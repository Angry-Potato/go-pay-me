package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Angry-Potato/go-pay-me/implementation/web"
)

func main() {
	serverPort := port(os.Getenv("PORT"), 8080)
	log.Fatal(web.StartServer(serverPort))
}

func port(envPort string, defaultPort int) int {
	if envPort == "" {
		return defaultPort
	}

	parsedPort, _ := strconv.Atoi(envPort)
	return parsedPort
}
