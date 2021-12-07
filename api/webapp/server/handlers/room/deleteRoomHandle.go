package room

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteRoomHandle deletes Room by ID
func DeleteRoomHandle() httprouter.Handle {

	rooms := dto.GetHotelRoomsDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, r := range rooms {
			if r.HotelRoomID == id { // delete object imitation =)
				rooms[index].HotelRoomID = 0
				json.NewEncoder(w).Encode(rooms)
				return
			}
		}

		http.Error(w, "Cant find Room", http.StatusBadRequest)
	}
}
