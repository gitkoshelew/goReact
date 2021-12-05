package hotel

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetHotelHandle returns Hotel by ID
func GetHotelHandle() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		json.NewEncoder(w).Encode(dto.HotelToDto(entity.GetHotelByID(id)))
	}
}
