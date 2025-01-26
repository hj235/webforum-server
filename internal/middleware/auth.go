package middleware

import (
	"net/http"

	"github.com/hj235/cvwo/internal/auth"
)

var Verifier func(http.Handler) http.Handler = auth.Verifier
var Authenticator func(http.Handler) http.Handler = auth.Authenticator
