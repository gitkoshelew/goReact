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

		req := &UserRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		result, err := db.Exec("UPDATE USERS set email = $1, password = $2, role = $3, verified = $4, first_name = $5, surname = $6, middle_name = $7, sex = $8, date_of_birth = $9, address = $10, phone = $11, photo = $12 WHERE id = $13",
			req.Email, req.Password, req.Role, req.Verified, req.Name, req.Surname, req.MiddleName, req.Sex, req.DateOfBirth, req.Address, req.Phone, req.Photo, req.UserID)

		if err != nil {
			panic(err)
		}
		log.Println(result.RowsAffected())

		w.WriteHeader(http.StatusCreated)
	}
}
