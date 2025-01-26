package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth
var Verifier func(http.Handler) http.Handler
var Authenticator func(http.Handler) http.Handler

func InitAuth() {
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
	Verifier = jwtauth.Verifier(tokenAuth)
	Authenticator = jwtauth.Authenticator(tokenAuth)
}

func GenerateToken(username string) (string, error) {
	expDuration, _ := time.ParseDuration("24h")
	claims := map[string]any{
		"username": username,
		"exp":      jwtauth.ExpireIn(expDuration),
	}
	_, tokenStr, err := tokenAuth.Encode(claims)
	if err != nil {
		return "", fmt.Errorf("error encoding claims: %w", err)
	}
	return tokenStr, nil
}

func RetrieveClaims(req *http.Request) (map[string]any, error) {
	_, claims, err := jwtauth.FromContext(req.Context())
	if err != nil {
		return nil, fmt.Errorf("error retrieving claims: %w", err)
	}
	return claims, nil
}
