package hotel

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetHotelByID ...
func GetHotelByID(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while parsing id. Err msg: %v", err)})
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while openong DB. Err msg: %v", err)})
			return
		}

		hotel, err := s.Hotel().FindByID(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while searching hotel. Err msg: %v", err)})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hotel)
	}
}
