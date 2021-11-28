package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"net/http"
	"strconv"
)

// HandleHotels opens a hotel page, URL: "/hotels". Shows all hotels, can search one by id
func HandleHotels() http.HandlerFunc {

	hotels := webapp.GetHotels()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hotels)
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

		if hotelFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(hotel)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
