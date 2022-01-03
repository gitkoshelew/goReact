package middleware

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// WhoAmI ...
func WhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		user := r.Context().Value(CtxKeyUser).(dto.UserDto)
		json.NewEncoder(w).Encode(user)
	}
}
