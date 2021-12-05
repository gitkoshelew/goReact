package employee

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PostEmployeesHandler creates Employee
func PostEmployeesHandler() http.HandlerFunc {

	employeesDto := dto.GetEmployeesDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &employeeRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		e := entity.Employee{
			User:       entity.GetUserByID(req.UserID),
			Hotel:      entity.GetHotelByID(req.HotelID),
			EmployeeID: req.EmployeeID,
			Position:   req.Position,
			Role:       req.Role,
		}
		employeesDto = append(employeesDto, dto.EmployeeDto(entity.EmployeeToDto(e)))
		json.NewEncoder(w).Encode(employeesDto)
	}
}
