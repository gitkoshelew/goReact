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

		req := &loginRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		user := store.User{}
		if err := db.QueryRow("SELECT * FROM users WHERE email = $1",
			req.Email).Scan(
			&user.UserID,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Verified,
			&user.Name,
			&user.Surname,
			&user.MiddleName,
			&user.Sex,
			&user.DateOfBirth,
			&user.Address,
			&user.Phone,
			&user.Photo,
		); err != nil {
			log.Print("email check failed")
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		err := store.CheckPasswordHash(user.Password, req.Password)
		if err != nil {
			log.Print("email check failed")
			http.Error(w, "Invalid login or password", http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		tk, err := CreateToken(uint64(user.UserID))
		tokens := map[string]string{
			"access_token":  tk.AccessToken,
			"refresh_token": tk.RefreshToken,
		}
		// err = CreateAuth(uint64(user.UserID), tk)
		// if err != nil {
		// 	log.Printf("%v. %v", http.StatusUnprocessableEntity, err)
		// }

		c := http.Cookie{
			Name:     "JWT",
			Value:    tk.AccessToken,
			HttpOnly: true,
		}

		http.SetCookie(w, &c)
		json.NewEncoder(w).Encode(tokens)
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusCreated)
	}
}
