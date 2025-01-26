package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/hj235/cvwo/internal/handlers/users"
	"github.com/hj235/cvwo/internal/middleware"
)

func GetUserRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		router.Use(middleware.DefaultMiddleware)

		router.Group(publicUserRoutes())
		router.Group(protectedUserRoutes())
	}
}

func publicUserRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		router.Group(func(router chi.Router) {
			router.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, _ := users.HandleListAll(w, req)
				json.NewEncoder(w).Encode(response)
			})
		})

		router.Post("/signup", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleSignup(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Post("/login", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleLogin(w, req)
			json.NewEncoder(w).Encode(response)
		})
	}
}

func protectedUserRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		// router.Use(middleware.Verifier)
		// router.Use(middleware.Authenticator)
		router.Use(middleware.GetCorsMiddleware())

		router.Patch("/edit/{username}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleEdit(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Patch("/delete", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleDelete(w, req)
			json.NewEncoder(w).Encode(response)
		})
	}
}
