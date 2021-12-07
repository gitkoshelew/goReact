package employee

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PutEmployeesHandler update Employee
func PutEmployeesHandler() http.HandlerFunc {

	employeesDto := dto.GetEmployeesDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &employeeRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, e := range employeesDto {
			if e.EmployeeID == req.EmployeeID {
				employeesDto[index].HotelID = req.HotelID
				employeesDto[index].Position = req.Position
				employeesDto[index].Role = req.Role
				break
			}
		}

		json.NewEncoder(w).Encode(employeesDto)
	}
}
