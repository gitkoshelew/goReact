package room

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// GetRoomsHandler returns all Rooms
func GetRoomsHandler() http.HandlerFunc {
	rooms := dto.GetHotelRoomsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rooms)
	}
}
