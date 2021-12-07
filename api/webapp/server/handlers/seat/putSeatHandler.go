package seat

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PutSeatHandler updates Seat
func PutSeatHandler() http.HandlerFunc {

	seats := dto.GetHotelRoomSeatsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &seatRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, s := range seats {
			if s.HotelRoomSeatID == req.HotelRoomSeatID {
				seats[index].Description = req.Description
				seats[index].IsFree = req.IsFree
				break
			}
		}

		json.NewEncoder(w).Encode(seats)
	}
}
