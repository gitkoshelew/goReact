package hotelhandlers

import (
	"encoding/json"
	"fmt"
	"hotel/domain/store"
	"hotel/internal/apperror"
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
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id")),
				fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))))
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		hotel, err := s.Hotel().FindByID(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting hotel by id.",
				fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Error occured while getting hotel by id. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hotel)
	}
}
