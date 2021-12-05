package user

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// GetUsersHandler returns all Users
func GetUsersHandler() http.HandlerFunc {
	users := dto.GetUsersDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
