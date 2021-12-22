package employee

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetEmployeeHandle returns Employee by ID
func GetEmployeeHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		row := db.QueryRow("SELECT * FROM EMPLOYEE WHERE id = $1", id)

		employee := dto.EmployeeDto{}
		err = row.Scan(
			&employee.EmployeeID,
			&employee.UserID,
			&employee.HotelID,
			&employee.Position,
			&employee.Role)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(employee)
	}
}
