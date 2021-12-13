package hotel

import (
	"encoding/json"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PutHotelHandle updates Hotel
func PutHotelHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &hotelRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		result, err := db.Exec("UPDATE HOTEL set name = $1, address = $2 WHERE id = $3",
			req.Name, req.Address, req.HotelID)

		if err != nil {
			panic(err)
		}
		log.Println(result.RowsAffected())

		w.WriteHeader(http.StatusCreated)
	}
}
