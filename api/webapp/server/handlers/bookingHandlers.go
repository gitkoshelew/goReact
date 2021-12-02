package handlers

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/pkg/date"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type bookingRequest struct {
	BookingID   int       `json:"bookingId"`
	PetID       int       `json:"pet"`
	SeatID      int       `json:"seat"`
	Status      string    `json:"status"`
	StartDate   date.Date `json:"start"`
	EndDate     date.Date `json:"end"`
	EmployeeID  int       `json:"employeeId"`
	ClientNotes string    `json:"notes"`
}

// HandleBookings GET /api/bookings - returns all booking(JSON)
//		   	      POST /api/booking - add booking(JSON)
//	    	   	  PUT /api/booking - update booking(JSON)
func HandleBookings() http.HandlerFunc {

	bookings := entity.GetBookings()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bookings)
		// POST
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &bookingRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			b := entity.Booking{
				BookingID:   req.BookingID,
				Pet:         entity.GetPetByID(req.PetID),
				Seat:        entity.GetSeatByID(req.SeatID),
				Status:      entity.BookingStatus(req.Status),
				StartDate:   req.StartDate,
				EndDate:     req.EndDate,
				Employee:    entity.GetEmployeeByID(req.EmployeeID),
				ClientNotes: req.ClientNotes,
			}
			bookings = append(bookings, b)
			json.NewEncoder(w).Encode(bookings)
		// PUT
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &bookingRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, b := range bookings {
				if b.BookingID == req.BookingID {
					bookings[index].Pet = entity.GetPetByID(req.BookingID)
					bookings[index].Seat = entity.GetSeatByID(req.PetID)
					bookings[index].Status = entity.BookingStatus(req.Status)
					bookings[index].StartDate = req.StartDate
					bookings[index].EndDate = req.EndDate
					bookings[index].Employee = entity.GetEmployeeByID(req.EmployeeID)
					bookings[index].ClientNotes = req.ClientNotes
					break
				}
			}
			json.NewEncoder(w).Encode(bookings)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleBooking GET /api/booking/:id - returns booking by ID (JSON)
// 		 		 DELETE /api/booking/:id - delete booking by ID(JSON)
func HandleBooking() httprouter.Handle {

	bookings := entity.GetBookings()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}
			json.NewEncoder(w).Encode(entity.GetBookingByID(id))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, b := range bookings {
				if b.BookingID == id { // delete object imitation =)
					bookings[index].ClientNotes = "DELETE"
					json.NewEncoder(w).Encode(bookings)
					return
				}
			}

			http.Error(w, "Cant find Booking", http.StatusBadRequest)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
