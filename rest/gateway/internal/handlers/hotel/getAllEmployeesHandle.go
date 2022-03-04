package hotel

import (
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/hotel"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAllEmployeesHandle ...
func GetAllEmployeesHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		getAllService, err := hotel.GetAll(r.Context(), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var employees []*EmployeeDTO
		if err := json.Unmarshal(getAllService.Body, &employees); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(employees)
	}
}
