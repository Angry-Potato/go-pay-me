package web

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// StartServer starts the server
func StartServer() error {
	router, err := rest.MakeRouter()

	if err != nil {
		return err
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(router)

	return http.ListenAndServe(":8080", api.MakeHandler())
}
