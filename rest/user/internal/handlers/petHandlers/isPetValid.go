package pethandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user/internal/apperror"
	"user/internal/store"

	"github.com/julienschmidt/httprouter"
)

// PetValidationHandle ...
func PetValidationHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("petId")), fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))))
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't open DB. Err msg:%v.", err)))
			return
		}

		isValid, err := s.PetRepository.IsPetValid(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("err occured while requesting DB", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("err occured while requesting DB. Err msg: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(isValid)
	}
}
