package middleware

import (
	"fmt"
	"goReact/webapp/server/handlers/authentication"
	"log"
	"net/http"
)

// IsLoggedIn checks if user is loggen in, if true - return, if false - execute login handle
func IsLoggedIn(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := authentication.ExtractTokenMetadata(r)
		if err != nil {
			fmt.Fprintln(w, "You are not uthorized")
			log.Print(err.Error())
			next.ServeHTTP(w, r)
			return
		}
		fmt.Fprintln(w, "You are already logged in")
		return

	})
}
