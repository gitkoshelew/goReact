package seat

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetSeatHandle returns Seat by ID
func GetSeatHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		row := db.QueryRow("SELECT * FROM SEAT WHERE id = $1", id)

		seat := dto.SeatDto{}
		err = row.Scan(
			&seat.SeatID,
			&seat.RoomID,
			&seat.IsFree,
			&seat.Description)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(seat)
	}
}
