package employee

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetEmployeesHandle returns all Employees
func GetEmployeesHandle() httprouter.Handle {

	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT * FROM EMPLOYEE")
		if err != nil {
			log.Fatal(err)
		}

		employeesDto := []dto.EmployeeDto{}

		for rows.Next() {
			employee := dto.EmployeeDto{}
			err := rows.Scan(
				&employee.EmployeeID,
				&employee.UserID,
				&employee.HotelID,
				&employee.Position,
				&employee.Role)

			if err != nil {
				log.Printf(err.Error())
				continue
			}

			employeesDto = append(employeesDto, employee)
		}

		json.NewEncoder(w).Encode(employeesDto)
	}
}
