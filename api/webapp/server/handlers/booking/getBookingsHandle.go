package booking

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetBookingsHandle returns all bookings
func GetBookingsHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT * FROM BOOKING")
		if err != nil {
			log.Fatal(err)
		}
		bookingssDto := []dto.BookingDto{}

		for rows.Next() {
			booking := dto.BookingDto{}
			err := rows.Scan(
				&booking.BookingID,
				&booking.SeatID,
				&booking.PetID,
				&booking.EmployeeID,
				&booking.Status,
				&booking.StartDate,
				&booking.EndDate,
				&booking.ClientNotes)
			if err != nil {
				log.Printf(err.Error())
				continue
			}
			bookingssDto = append(bookingssDto, booking)
		}
		json.NewEncoder(w).Encode(bookingssDto)
	}
}
