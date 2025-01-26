package middleware

import (
	"net/http"
	"os"

	"github.com/go-chi/cors"
)

func GetCorsMiddleware() func(next http.Handler) http.Handler {
	var corsOptions = cors.Options{
		AllowedOrigins:   []string{os.Getenv("CLIENT_URL")},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}
	return cors.Handler(corsOptions)
}
