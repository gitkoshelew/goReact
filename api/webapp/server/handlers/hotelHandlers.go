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

// HandleHotels opens a hotel page, URL: "/hotels". Shows all hotels, can search one by id
func HandleHotels() http.HandlerFunc {

	hotels := webapp.GetHotels()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/hotels.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "hotels", hotels)
	}
}

// HandleHotelSearch shows a hotel by id, URL"/hotel?id="
func HandleHotelSearch() http.HandlerFunc {

	hotels := webapp.GetHotels()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var hotel entity.Hotel
		hotelFound := false

		for _, a := range hotels {
			if a.HotelID == id {
				hotel = a
				hotelFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_hotel.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if hotelFound {
			tmpl.ExecuteTemplate(w, "show_hotel", hotel)
		} else {
			tmpl.ExecuteTemplate(w, "show_hotel", "Hotel not found")
		}
	}
}
