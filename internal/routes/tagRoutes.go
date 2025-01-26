package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/hj235/cvwo/internal/handlers/users"
	"github.com/hj235/cvwo/internal/middleware"
)

func GetTagRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		router.Use(middleware.DefaultMiddleware)

		router.Group(publicTagRoutes())
		router.Group(protectedTagRoutes())
	}
}

func publicTagRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleListAll(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleSignup(w, req)
			json.NewEncoder(w).Encode(response)
		})
	}
}

func protectedTagRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		// router.Use(middleware.Verifier)
		// router.Use(middleware.Authenticator)
		router.Use(middleware.GetCorsMiddleware())

		router.Post("/create", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleLogin(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Patch("/edit/{id}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleEdit(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Delete("/delete/{id}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleDelete(w, req)
			json.NewEncoder(w).Encode(response)
		})
	}
}
