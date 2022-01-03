package middleware

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// EmployeetHome ...
func EmployeetHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		user := r.Context().Value(CtxKeyUser).(dto.UserDto)

		if store.Role(user.Role) != store.EmployeeRole {
			http.Error(w, "Sorry, you are not an employee", http.StatusUnauthorized)
			return
		}
		fmt.Print("Welcome employee!")
		json.NewEncoder(w).Encode(user)
	}
}
