package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/hj235/cvwo/internal/handlers/threads"
	"github.com/hj235/cvwo/internal/middleware"
)

func GetThreadRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		router.Use(middleware.DefaultMiddleware)

		router.Group(publicThreadRoutes())
		router.Group(protectedThreadRoutes())
	}
}

func publicThreadRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, req *http.Request) {
			response, _ := threads.HandleList(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := threads.HandleGet(w, req)
			json.NewEncoder(w).Encode(response)
		})
	}
}

func protectedThreadRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		// router.Use(middleware.Verifier)
		// router.Use(middleware.Authenticator)
		router.Use(middleware.GetCorsMiddleware())

		router.Post("/create/{username}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := threads.HandleCreate(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Patch("/edit/{id}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := threads.HandleEdit(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Delete("/delete/{id}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := threads.HandleDelete(w, req)
			json.NewEncoder(w).Encode(response)
		})
	}
}
