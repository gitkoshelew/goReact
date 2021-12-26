package authentication

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/handlers/user"
	"goReact/webapp/server/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegistrationHandle ...
func RegistrationHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &user.UserRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		encryptedPassword, err := store.EncryptPassword(req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		var emailIsUsed bool
		err = db.QueryRow("SELECT EXISTS (SELECT email FROM users WHERE email = $1)", req.Email).Scan(&emailIsUsed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if emailIsUsed {
			http.Error(w, "Email already in use!", http.StatusBadRequest)
			return
		}

		var id int
		err = db.QueryRow("INSERT INTO USERS (email, password, role, verified, first_name, surname, middle_name, sex, date_of_birth, address, phone, photo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id",
			req.Email, encryptedPassword, req.Role, req.Verified, req.Name, req.Surname, req.MiddleName, req.Sex, req.DateOfBirth, req.Address, req.Phone, req.Photo,
		).Scan(&id)

		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(id)
		w.WriteHeader(http.StatusCreated)
	}
}
