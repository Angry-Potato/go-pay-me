package main

import (
	"log"

	"github.com/Angry-Potato/go-pay-me/implementation/web"
)

func main() {
	log.Fatal(web.StartServer(8080))
}
