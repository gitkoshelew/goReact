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

// HandleHotelRoomSeats opens an hotel room seats page, URL: "/seats". Shows all hotel room seats, can search one by id
func HandleHotelRoomSeats() http.HandlerFunc {

	hotelRoomSeats := webapp.GetHotelRoomSeats()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/seats.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "seats", hotelRoomSeats)
	}
}

// HandleHotelRoomSeatSearch shows an hotel room seats by id, URL"/seat?id="
func HandleHotelRoomSeatSearch() http.HandlerFunc {

	seats := webapp.GetHotelRoomSeats()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var seat entity.HotelRoomSeat
		seatFound := false

		for _, a := range seats {
			if a.HotelRoomSeatID == id {
				seat = a
				seatFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_seat.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if seatFound {
			tmpl.ExecuteTemplate(w, "show_seat", seat)
		} else {
			tmpl.ExecuteTemplate(w, "show_seat", "Seat not found")
		}
	}
}
