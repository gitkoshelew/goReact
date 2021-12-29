package authentication

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LoginHandle is GET method. Checkes login and password and returns accounts user if validation was passed
func LoginHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		utils.EnableCors(&w)

		req := &loginRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		user := store.User{}
		if err := db.QueryRow("SELECT * FROM users WHERE email = $1",
			req.Email).Scan(
			&user.UserID,
			&user.Name,
			&user.Surname,
			&user.MiddleName,
			&user.Email,
			&user.DateOfBirth,
			&user.Address,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Verified,
			&user.Sex,
			&user.Photo,
		); err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		err := store.CheckPasswordHash(user.Password, req.Password)
		if err != nil {
			http.Error(w, "Invalid login or password", http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		tk, err := CreateToken(uint64(user.UserID))

		c := http.Cookie{
			Name:     "Refresh-Token",
			Value:    tk.RefreshToken,
			HttpOnly: true,
		}

		http.SetCookie(w, &c)
		w.Header().Add("Access-Token", tk.AccessToken)
		json.NewEncoder(w).Encode(user)
	}
}
