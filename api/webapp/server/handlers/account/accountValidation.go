package account

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Validation is GET method. Checkes login and password and returns accounts user if validation was passed
func Validation() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &accountRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		account := store.Account{}
		if err := db.QueryRow("SELECT * FROM ACCOUNT WHERE login = $1",
			req.Login).Scan(&account.AccountID, &account.Login, &account.Password); err != nil {
			http.Error(w, "Invalid login or password1", http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		err := store.CheckPasswordHash(account.Password, req.Password)
		if err != nil {
			http.Error(w, "Invalid login or password2", http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		row := db.QueryRow("SELECT * FROM USERS WHERE account_id = $1", account.AccountID)

		user := dto.UserDto{}
		err = row.Scan(
			&user.UserID,
			&user.Name,
			&user.Surname,
			&user.MiddleName,
			&user.Email,
			&user.DateOfBirth,
			&user.Address,
			&user.Phone,
			&user.AccountID)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(user)
	}
}
