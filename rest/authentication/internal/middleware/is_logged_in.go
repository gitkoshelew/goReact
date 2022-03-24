package middleware

import (
	"auth/internal/apperror"
	"auth/pkg/jwt"
	"encoding/json"
	"fmt"
	"net/http"
)

// IsLoggedIn checks if user is loggen in, if true - return, if false - execute login handle
func IsLoggedIn(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := jwt.ExtractTokenMetadata(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		json.NewEncoder(w).Encode(apperror.NewAppError("You are already logged in", fmt.Sprintf("%d", http.StatusOK), ""))
		return
	})
}
