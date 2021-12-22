package hotel

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetHotelsHandle returns all Hotels
func GetHotelsHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT * FROM HOTEL")
		if err != nil {
			log.Fatal(err)
		}

		hotelsDto := []dto.HotelDto{}

		for rows.Next() {
			hotel := dto.HotelDto{}
			err := rows.Scan(
				&hotel.HotelID,
				&hotel.Name,
				&hotel.Address)

			if err != nil {
				log.Printf(err.Error())
				continue
			}

			hotelsDto = append(hotelsDto, hotel)
		}

		json.NewEncoder(w).Encode(hotelsDto)
	}
}
