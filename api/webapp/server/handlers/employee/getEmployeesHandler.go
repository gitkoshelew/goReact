package employee

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// GetEmployeesHandler returns all Employees
func GetEmployeesHandler() http.HandlerFunc {

	employeesDto := dto.GetEmployeesDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(employeesDto)
	}
}
