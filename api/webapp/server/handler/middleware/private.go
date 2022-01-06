package middleware

import (
	"goReact/domain/store"
	"net/http"
)

// Private ...
func Private(next http.HandlerFunc, s *store.Store) http.Handler {
	return AuthenticateUser(next, s)
}
