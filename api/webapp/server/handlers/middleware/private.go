package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Private ...
func Private(next http.HandlerFunc) httprouter.Handle {
	return AuthenticateUser(next)
}
