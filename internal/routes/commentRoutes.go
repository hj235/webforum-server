package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/hj235/cvwo/internal/handlers/comments"
	"github.com/hj235/cvwo/internal/middleware"
)

func GetCommentRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		router.Use(middleware.DefaultMiddleware)

		router.Group(publicCommentRoutes())
		router.Group(protectedCommentRoutes())
	}
}

func publicCommentRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		router.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := comments.HandleList(w, req)
			json.NewEncoder(w).Encode(response)
		})
	}
}

func protectedCommentRoutes() func(router chi.Router) {
	return func(router chi.Router) {
		// router.Use(middleware.Verifier)
		// router.Use(middleware.Authenticator)
		router.Use(middleware.GetCorsMiddleware())

		router.Post("/create/{username}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := comments.HandleCreate(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Patch("/edit/{id}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := comments.HandleEdit(w, req)
			json.NewEncoder(w).Encode(response)
		})

		router.Delete("/delete/{id}", func(w http.ResponseWriter, req *http.Request) {
			response, _ := comments.HandleDelete(w, req)
			json.NewEncoder(w).Encode(response)
		})
	}
}
