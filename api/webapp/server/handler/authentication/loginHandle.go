package authentication

import (
	"encoding/json"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler/request"
	"goReact/webapp/server/handler/response"
	"net/http"
)

// LoginHandle checkes login and password and returns user if validation was passed
func LoginHandle(s *store.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := &request.Login{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
		}

		err := s.Open()
		if err != nil {
			s.Logger.Errorf("Can't open DB. Err msg: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		user, err := s.User().FindByEmail(req.Email)
		if err != nil {
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err = model.CheckPasswordHash(user.Password, req.Password)
		if err != nil {
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %s", err.Error())
			http.Error(w, "Invalid login or password", http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		tk, err := CreateToken(uint64(user.UserID), string(user.Role))
		if err != nil {
			s.Logger.Errorf("Eror during createing tokens. Err msg: %s", err.Error())
			http.Error(w, "Eror during createing tokens", http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

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
