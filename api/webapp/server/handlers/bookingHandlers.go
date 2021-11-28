package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"net/http"
	"strconv"
)

// HandleBookings opens an booking page, URL: "/bookings". Shows all bookings, can search one by id
func HandleBookings() http.HandlerFunc {

	bookings := webapp.GetBookings()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookings)
	}
}

// HandleBookingSearch shows an booking by id, URL"/booking?id="
func HandleBookingSearch() http.HandlerFunc {

	bookings := webapp.GetBookings()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var booking entity.Booking
		bookingFound := false

		for _, a := range bookings {
			if a.BookingID == id {
				booking = a
				bookingFound = true
				break
			}
		}

		if bookingFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(booking)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
