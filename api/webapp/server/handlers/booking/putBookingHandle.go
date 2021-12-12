package booking

import (
	"encoding/json"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PutBookingsHandle updates Booking
func PutBookingsHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &bookingRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		result, err := db.Exec("UPDATE BOOKING set seat_id = $1, pet_id = $2, employee_id = $3, status = $4, start_date =$5, end_date = $6, client_notes = $7 WHERE id = $8",
			req.SeatID, req.PetID, req.EmployeeID, req.Status, req.StartDate, req.EndDate, req.ClientNotes, req.BookingID)

		if err != nil {
			panic(err)
		}
		log.Println(result.RowsAffected())

		w.WriteHeader(http.StatusCreated)
	}
}
