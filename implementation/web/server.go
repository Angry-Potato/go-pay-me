package web

import (
	"fmt"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// StartServer starts the server
func StartServer(port int) error {
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

	return http.ListenAndServe(serverPort(port), api.MakeHandler())
}

func serverPort(port int) string {
	return fmt.Sprintf(":%d", port)
}
