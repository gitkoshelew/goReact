package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"net/http"
	"strconv"
)

// HandleHotelRooms opens an hotel rooms page, URL: "/rooms". Shows all hotel rooms, can search one by id
func HandleHotelRooms() http.HandlerFunc {

	rooms := webapp.GetHotelRooms()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rooms)
	}
}

// HandleHotelRoomSearch shows a hotel room by id, URL"/room?id="
func HandleHotelRoomSearch() http.HandlerFunc {

	rooms := webapp.GetHotelRooms()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var room entity.HotelRoom
		roomFound := false

		for _, a := range rooms {
			if a.HotelRoomID == id {
				room = a
				roomFound = true
				break
			}
		}

		if roomFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(room)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
