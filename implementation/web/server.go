package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
)

// StartServer starts the server
func StartServer(port int, DB *gorm.DB) error {
	api := rest.NewApi()
	statusMw := &rest.StatusMiddleware{}
	api.Use(statusMw)
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/.status", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(statusMw.GetStatus())
		}),
		rest.Get("/payments", AllPayments(DB)),
		rest.Post("/payments", CreatePayment(DB)),
		rest.Delete("/payments", DeleteAllPayments(DB)),
		rest.Put("/payments", SetAllPayments(DB)),
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
