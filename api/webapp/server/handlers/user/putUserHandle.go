package user

import (
	"encoding/json"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PutUserHandle updates User
func PutUserHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &userRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		result, err := db.Exec("UPDATE USERS set first_name = $1, surname = $2, middle_name = $3, email = $4, date_of_birth =$5, address = $6, phone = $7 WHERE id = $8",
			req.Name, req.Surname, req.MiddleName, req.Email, req.DateOfBirth, req.Address, req.Phone, req.UserID)

		if err != nil {
			panic(err)
		}
		log.Println(result.RowsAffected())

		w.WriteHeader(http.StatusCreated)
	}
}
