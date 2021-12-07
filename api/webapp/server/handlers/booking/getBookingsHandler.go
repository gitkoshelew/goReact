package booking

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// GetBookingsHandler returns all bookings
func GetBookingsHandler() http.HandlerFunc {
	bookingsDto := dto.GetBookingsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookingsDto)
	}
}
