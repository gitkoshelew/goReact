package middleware

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// ClientHome ...
func ClientHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		user := r.Context().Value(CtxKeyUser).(dto.UserDto)

		if store.Role(user.Role) != store.ClientRole {
			http.Error(w, "Sorry, you are not a client", http.StatusUnauthorized)
			return
		}
		fmt.Print("Welcome Client!")
		json.NewEncoder(w).Encode(user)
	}
}
