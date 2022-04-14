package booking

import (
	"encoding/json"
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"net/http"
)

// PostBookingHandle ...
func 
PostBookingHandle(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		bookingDTO := r.Context().Value(handler.CtxKeyBookingValidation).(*model.BookingDTO)

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error occured while opening DB, err: %v", err)})
			return
		}

		booking, err := s.Booking().ModelFromDTO(bookingDTO)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error occured while creating booking, err: %v", err)})
			return
		}

		_, err = s.Booking().Create(booking)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error occured while creating booking, err: %v", err)})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Booking id = %d", booking.BookingID)})
	}
}
