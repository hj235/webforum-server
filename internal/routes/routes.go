package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("Hello World!"))
		})
	}
}
