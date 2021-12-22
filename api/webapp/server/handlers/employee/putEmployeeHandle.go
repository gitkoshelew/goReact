package employee

import (
	"encoding/json"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PutEmployeesHandle update Employee
func PutEmployeesHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &employeeRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		result, err := db.Exec("UPDATE EMPLOYEE set user_id = $1, hotel_id = $2, position = $3, role = $4 WHERE id = $5",
			req.UserID, req.HotelID, req.Position, req.Role, req.EmployeeID)

		if err != nil {
			panic(err)
		}

		log.Println(result.RowsAffected())

		w.WriteHeader(http.StatusCreated)
	}
}
