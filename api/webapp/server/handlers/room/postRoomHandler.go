package room

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PostRoomHandler creates Room
func PostRoomHandler() http.HandlerFunc {
	rooms := dto.GetHotelRoomsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &roomRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		room := entity.HotelRoom{
			HotelRoomID: req.HotelRoomID,
			RoomNumber:  req.RoomNumber,
			PetType:     entity.PetType(req.PetType),
			Seats:       entity.GetSeatsByID(req.SeatsID),
		}

		rooms = append(rooms, dto.RoomDto(entity.RoomToDto(room)))
		json.NewEncoder(w).Encode(rooms)
	}
}
