package authentication

import (
	"encoding/json"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler/request"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegistrationHandle ...
func RegistrationHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &request.User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			http.Error(w, "Bad request", http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}

		u := model.NewUser(
			0,
			req.Email,
			req.Password,
			req.Role,
			req.Name,
			req.Surname,
			req.MiddleName,
			req.Sex,
			req.Address,
			req.Phone,
			req.Photo,
			req.Verified,
			req.DateOfBirth,
		)

		err := s.Open()
		if err != nil {
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		_, err = s.User().Create(&u)
		if err != nil {
			s.Logger.Errorf("Cant create user. Err msg: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode(u.UserID)
		w.WriteHeader(http.StatusCreated)
	}
}
