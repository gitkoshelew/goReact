package seat

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteSeatHandle deletes Room by ID
func DeleteSeatHandle() httprouter.Handle {
	seats := dto.GetHotelRoomSeatsDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, s := range seats {
			if s.HotelRoomSeatID == id { // delete object imitation =)
				seats[index].Description = "DELETE"
				json.NewEncoder(w).Encode(seats)
				return
			}
		}

		http.Error(w, "Cant find Seat", http.StatusBadRequest)
	}
}
