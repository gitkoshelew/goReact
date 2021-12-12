package user

import (
	"encoding/json"
	"goReact/webapp/server/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PostUserHandle creates User
func PostUserHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &userRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		var id int
		err := db.QueryRow("INSERT into USERS (first_name, surname, middle_name, email, date_of_birth, address, phone, account_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
			req.Name, req.Surname, req.MiddleName, req.Email, req.DateOfBirth, req.Address, req.Phone, req.AccountID,
		).Scan(&id)

		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(id)
		w.WriteHeader(http.StatusCreated)
	}
}
