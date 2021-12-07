package seat

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// GetSeatsHandler returns all Seats
func GetSeatsHandler() http.HandlerFunc {
	seats := dto.GetHotelRoomSeatsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(seats)
	}
}
