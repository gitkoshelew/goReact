package authentication

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LoginHandle is GET method. Checkes login and password and returns accounts user if validation was passed
func LoginHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()
	type accountRequest struct {
		AccountID int    `json:"accountId"`
		Login     string `json:"login"`
		Password  string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &accountRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		account := store.Account{}
		if err := db.QueryRow("SELECT * FROM ACCOUNT WHERE login = $1",
			req.Login).Scan(&account.AccountID, &account.Login, &account.Password); err != nil {
			http.Error(w, "Invalid login or password", http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		err := store.CheckPasswordHash(account.Password, req.Password)
		if err != nil {
			http.Error(w, "Invalid login or password", http.StatusUnauthorized)
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

		tk, err := CreateToken(uint64(user.AccountID))
		tokens := map[string]string{
			"access_token":  tk.AccessToken,
			"refresh_token": tk.RefreshToken,
		}
		err = CreateAuth(uint64(user.AccountID), tk)
		if err != nil {
			log.Printf("%v. %v", http.StatusUnprocessableEntity, err)
		}

		c := http.Cookie{
			Name:     "JWT",
			Value:    tk.AccessToken,
			HttpOnly: true,
		}

		http.SetCookie(w, &c)
		json.NewEncoder(w).Encode(tokens)
		json.NewEncoder(w).Encode(user)

	}
}
