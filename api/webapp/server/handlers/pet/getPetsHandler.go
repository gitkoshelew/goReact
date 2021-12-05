package pet

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// GetPetsHandler returns all Pets
func GetPetsHandler() http.HandlerFunc {
	pets := dto.GetPetsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pets)
	}
}
