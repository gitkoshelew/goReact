package middleware

import (
	"fmt"
	"goReact/webapp/server/handler/authentication"
	"net/http"
)

// IsLoggedIn checks if user is loggen in, if true - return, if false - execute login handle
func IsLoggedIn(next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := authentication.ExtractTokenMetadata(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Fprintln(w, "You are already logged in")
		return

	})
}
