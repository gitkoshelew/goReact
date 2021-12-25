package booking

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetBookingHandle returns Booking by ID
func GetBookingHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		row := db.QueryRow("SELECT * FROM BOOKING WHERE id = $1", id)

		booking := dto.BookingDto{}
		err = row.Scan(
			&booking.BookingID,
			&booking.SeatID,
			&booking.PetID,
			&booking.EmployeeID,
			&booking.Status,
			&booking.StartDate,
			&booking.EndDate,
			&booking.ClientNotes)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(booking)
	}
}
