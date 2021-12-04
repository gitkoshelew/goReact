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
	PetID       int       `json:"petId"`
	SeatID      int       `json:"seatId"`
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

	bookingsDto := entity.GetBookingsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bookingsDto)
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
			bookingsDto = append(bookingsDto, entity.BookingToDto(b))
			json.NewEncoder(w).Encode(bookingsDto)
		// PUT
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &bookingRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, b := range bookingsDto {
				if b.BookingID == req.BookingID {
					bookingsDto[index].PetID = req.PetID
					bookingsDto[index].SeatID = req.SeatID
					bookingsDto[index].Status = req.Status
					bookingsDto[index].StartDate = req.StartDate
					bookingsDto[index].EndDate = req.EndDate
					bookingsDto[index].EmployeeID = req.EmployeeID
					bookingsDto[index].ClientNotes = req.ClientNotes
					break
				}
			}
			json.NewEncoder(w).Encode(bookingsDto)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleBooking GET /api/booking/:id - returns booking by ID (JSON)
// 		 		 DELETE /api/booking/:id - delete booking by ID(JSON)
func HandleBooking() httprouter.Handle {

	bookingsDto := entity.GetBookingsDto()

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
			json.NewEncoder(w).Encode(entity.BookingToDto(entity.GetBookingByID(id)))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, b := range bookingsDto {
				if b.BookingID == id { // delete object imitation =)
					bookingsDto[index].ClientNotes = "DELETE"
					json.NewEncoder(w).Encode(bookingsDto)
					return
				}
			}

			http.Error(w, "Cant find Booking", http.StatusBadRequest)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
