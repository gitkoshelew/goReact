package pethandlers

import (
	"customer/internal/apperror"
	"customer/internal/store"
	"customer/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeletePet ...
func DeletePet(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

		err = s.Pet().Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while deleting pet ", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Deleted pet with id = %d", id)})

	}
}
