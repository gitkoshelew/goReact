package booking

import (
	"encoding/json"
	"goReact/webapp/server/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PostBookingsHandle creates Booking
func PostBookingsHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &bookingRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		var id int
		err := db.QueryRow("INSERT into BOOKING (seat_id , pet_id, employee_id, status, start_date, end_date, client_notes) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
			req.SeatID, req.PetID, req.EmployeeID, req.Status, req.StartDate, req.EndDate, req.ClientNotes,
		).Scan(&id)

		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(id)
		w.WriteHeader(http.StatusCreated)
	}
}
