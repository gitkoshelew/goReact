package handlers

import (
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// HandleBookings opens an booking page, URL: "/bookings". Shows all bookings, can search one by id
func HandleBookings() http.HandlerFunc {

	bookings := webapp.GetBookings()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/bookings.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "bookings", bookings)
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

		tmpl, err := template.ParseFiles("webapp/templates/show_booking.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if bookingFound {
			tmpl.ExecuteTemplate(w, "show_booking", booking)
		} else {
			tmpl.ExecuteTemplate(w, "show_booking", "Booking not found")
		}
	}
}
