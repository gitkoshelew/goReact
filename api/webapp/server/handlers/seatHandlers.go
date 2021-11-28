package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"net/http"
	"strconv"
)

// HandleHotelRoomSeats opens an hotel room seats page, URL: "/seats". Shows all hotel room seats, can search one by id
func HandleHotelRoomSeats() http.HandlerFunc {

	seats := webapp.GetHotelRoomSeats()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(seats)
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

		if seatFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(seat)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
