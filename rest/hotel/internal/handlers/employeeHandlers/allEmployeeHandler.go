package employeehandlers

import (
	"encoding/json"
	"fmt"
	"hotel/internal/apperror"
	"hotel/internal/store"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AllEmployeeHandler ...
func AllEmployeeHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't open DB. Err msg:%v.", err)))
			return
		}

		employees, err := s.Employee().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't find employees", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't find employees. Err msg: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(employees)

	}
}
