package employee

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteEmployeeHandle deletes Employee by ID
func DeleteEmployeeHandle() httprouter.Handle {
	employeesDto := dto.GetEmployeesDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, e := range employeesDto {
			if e.EmployeeID == id { // delete object imitation =)
				employeesDto[index].Position = "DELETE"
				employeesDto[index].Role = "DELETE"
				json.NewEncoder(w).Encode(employeesDto)
				return
			}
		}

		http.Error(w, "Cant find Employee", http.StatusBadRequest)
	}
}
