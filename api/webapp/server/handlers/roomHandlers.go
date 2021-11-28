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

// HandleHotelRooms opens an hotel rooms page, URL: "/rooms". Shows all hotel rooms, can search one by id
func HandleHotelRooms() http.HandlerFunc {

	rooms := webapp.GetHotelRooms()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/rooms.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "rooms", rooms)
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

		tmpl, err := template.ParseFiles("webapp/templates/show_room.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if roomFound {
			tmpl.ExecuteTemplate(w, "show_room", room)
		} else {
			tmpl.ExecuteTemplate(w, "show_room", "Room not found")
		}
	}
}
