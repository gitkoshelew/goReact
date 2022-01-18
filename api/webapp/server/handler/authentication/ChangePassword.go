package authentication

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/request"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// New pass ...
func ChangePassword(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &request.Login{}
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

		user, err := s.User().FindByEmail(req.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		fmt.Print("//1user.Password reqest  -", req.Password)//1
		fmt.Print("   ///2user.Password 2  -", req.Password)//2
		user.Password = req.Password
		/*err = user.WithEncryptedPassword()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		fmt.Println("   ///3user.Password ENCRYPTED-  ", user.Password)//3*/

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
