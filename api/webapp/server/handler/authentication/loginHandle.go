package authentication

import (
	"encoding/json"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler/request"
	"net/http"
)

// LoginHandle checkes login and password and returns user if validation was passed
func LoginHandle(s *store.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		req := &request.Login{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Logger.Errorf("Eror during JSON request decoding. Err msg: %s", err.Error())
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		err := s.Open()
		if err != nil {
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		user, err := s.User().FindByEmail(req.Email)
		if err != nil {
			s.Logger.Errorf("Eror during searching user by email. Err msg: %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = model.CheckPasswordHash(user.Password, req.Password)
		if err != nil {
			s.Logger.Errorf("Eror during checking users password. Err msg: %s", err.Error())
			http.Error(w, "Invalid login or password", http.StatusBadRequest)
			return
		}

		tk, err := CreateToken(uint64(user.UserID), string(user.Role))

		c := http.Cookie{
			Name:     "Refresh-Token",
			Value:    tk.RefreshToken,
			HttpOnly: true,
		}

		http.SetCookie(w, &c)

		w.Header().Add("Access-Token", tk.AccessToken)
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusOK)

	})
}
