package web

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Angry-Potato/go-pay-me/implementation/db"
	"github.com/ant0ine/go-json-rest/rest"
)

// StartServer starts the server
func StartServer(port int) error {
	_, err := db.Connect(
		os.Getenv("DATABASE_URL"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		return err
	}

	api := rest.NewApi()
	statusMw := &rest.StatusMiddleware{}
	api.Use(statusMw)
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/.status", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(statusMw.GetStatus())
		}),
	)

	if err != nil {
		return err
	}

	api.SetApp(router)

	listenPort := serverPort(port)
	log.Printf("Listening on port %s", listenPort)
	return http.ListenAndServe(listenPort, api.MakeHandler())
}

func serverPort(port int) string {
	return fmt.Sprintf(":%d", port)
}
