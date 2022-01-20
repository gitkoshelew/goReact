package restorePassword

import (
	"encoding/json"
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler/middleware"
	"goReact/webapp/server/handler/request"
	"goReact/webapp/server/handler/response"
	"net/http"
)

// New pass ...
func ChangePassword(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		req := &request.Login{}

		email := middleware.ContextEmail(r.Context())

		fmt.Println("EMAIL FORM CONTEX ///*/*/*/*/*/*", email)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		user, err := s.User().FindByEmail(email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		fmt.Println(" user   ", user.UserID)
		fmt.Println(" req.Password and user.Password 455 ", req.Password, user.Password)

		err = model.CheckPasswordHash(user.Password, req.Password)
		fmt.Println(" ERR   ", err)
		if err == nil {
			fmt.Println(" ERR req.Password == user.Password  ", err)
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Password cannot be the same")
			json.NewEncoder(w).Encode(response.Error{Messsage: "Password cannot be the same"})
			return
		}

		user.Password = req.Password

		err = s.User().PasswordChange(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Cant update password. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Password changed for user= %d", user.UserID)})
	}
}
