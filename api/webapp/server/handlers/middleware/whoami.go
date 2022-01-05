package middleware

import (
	"encoding/json"
	"goReact/domain/model"
	"net/http"
)

// WhoAmI ...
func WhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		user := r.Context().Value(CtxKeyUser).(model.User)
		json.NewEncoder(w).Encode(user)
	}
}
