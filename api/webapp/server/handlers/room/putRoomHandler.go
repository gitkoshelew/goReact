package room

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PutRoomHandler updates Room
func PutRoomHandler() http.HandlerFunc {
	rooms := dto.GetHotelRoomsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &roomRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, r := range rooms {
			if r.HotelRoomID == req.HotelRoomID {
				rooms[index].PetType = req.PetType
				rooms[index].RoomNumber = req.RoomNumber
				rooms[index].SeatsID = req.SeatsID
				break
			}
		}

		json.NewEncoder(w).Encode(rooms)
	}
}
