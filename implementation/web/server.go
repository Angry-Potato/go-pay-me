package web

import (
	"fmt"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// StartServer starts the server
func StartServer(port int) error {
	router, err := rest.MakeRouter()

	if err != nil {
		return err
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(router)

	return http.ListenAndServe(serverPort(port), api.MakeHandler())
}

func serverPort(port int) string {
	return fmt.Sprintf(":%d", port)
}
