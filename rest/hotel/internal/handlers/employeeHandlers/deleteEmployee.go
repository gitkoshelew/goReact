package employeehandlers

import (
	"encoding/json"
	"fmt"
	"hotel/internal/apperror"
	"hotel/internal/store"
	"hotel/pkg/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteEmployee ...
func DeleteEmployee(s *store.Store) httprouter.Handle {
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
		err = s.Employee().Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while deleting employee.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Error occured while deleting employee. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Deleted employee with id = %d", id)})

	}
}
