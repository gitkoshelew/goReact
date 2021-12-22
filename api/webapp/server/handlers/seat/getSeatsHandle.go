package seat

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetSeatsHandle returns all Seats
func GetSeatsHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT * FROM SEAT")
		if err != nil {
			log.Fatal(err)
		}

		seatsDto := []dto.SeatDto{}

		for rows.Next() {
			seat := dto.SeatDto{}
			err := rows.Scan(
				&seat.SeatID,
				&seat.RoomID,
				&seat.IsFree,
				&seat.Description)

			if err != nil {
				log.Printf(err.Error())
				continue
			}

			seatsDto = append(seatsDto, seat)
		}

		json.NewEncoder(w).Encode(seatsDto)
	}
}
