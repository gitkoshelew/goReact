package hotel

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// GetHotelsHandler returns all Hotels
func GetHotelsHandler() http.HandlerFunc {
	hotels := dto.GetHotelsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hotels)
	}
}
