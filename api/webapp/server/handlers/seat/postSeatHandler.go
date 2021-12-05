package seat

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PostSeatHandler creates Seat
func PostSeatHandler() http.HandlerFunc {
	seats := dto.GetHotelRoomSeatsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &seatRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		s := entity.HotelRoomSeat{
			HotelRoomSeatID: req.HotelRoomSeatID,
			Description:     req.Description,
			IsFree:          req.IsFree,
		}

		seats = append(seats, dto.SeatToDto(s))
		json.NewEncoder(w).Encode(seats)
	}
}
