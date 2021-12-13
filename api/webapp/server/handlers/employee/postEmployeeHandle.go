package employee

import (
	"encoding/json"
	"goReact/webapp/server/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PostEmployeesHandle creates Employee
func PostEmployeesHandle() httprouter.Handle {

	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &employeeRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		var id int
		err := db.QueryRow("INSERT into EMPLOYEE (user_id, hotel_id, position, role) VALUES ($1, $2, $3, $4) RETURNING id",
			req.UserID, req.HotelID, req.Position, req.Role,
		).Scan(&id)

		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(id)
		w.WriteHeader(http.StatusCreated)
	}
}
