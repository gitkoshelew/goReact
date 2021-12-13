package hotel

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetHotelHandle returns Hotel by ID
func GetHotelHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(
		w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		row := db.QueryRow("SELECT * FROM HOTEL WHERE id = $1", id)

		hotel := dto.HotelDto{}
		err = row.Scan(
			&hotel.HotelID,
			&hotel.Name,
			&hotel.Address)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(hotel)
	}
}
