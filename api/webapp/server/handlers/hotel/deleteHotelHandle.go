package hotel

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteHotel deletes Hotel by ID
func DeleteHotel() httprouter.Handle {
	hotels := dto.GetHotelsDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, h := range hotels {
			if h.HotelID == id { // delete object imitation =)
				hotels[index].Name = "DELETE"
				json.NewEncoder(w).Encode(hotels)
				return
			}
		}

		http.Error(w, "Cant find Hotel", http.StatusBadRequest)
	}
}
