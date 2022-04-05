package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ValidateBooking ...
func ValidateBooking(next http.Handler, s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		bookingDTO := &model.BookingDTO{}
		if err := json.NewDecoder(r.Body).Decode(bookingDTO); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		if bookingDTO.EndDate.Before(*bookingDTO.StartDate) {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Error("End date can not be before start date")
			json.NewEncoder(w).Encode(response.Error{Messsage: "End date can not be before start date."})
			return
		}

		err := bookingDTO.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while validating booking. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while validating booking. Err msg: %v", err)})
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(context.Background(), handler.CtxKeyBookingValidation, bookingDTO)))
	}

}
