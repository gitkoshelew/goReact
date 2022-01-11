package user

import (
	"encoding/json"
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler/request"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PostUserHandle creates User
func PostUserHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &request.User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			http.Error(w, "Bad request", http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
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
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		_, err = s.User().Create(&u)
		if err != nil {
			s.Logger.Errorf("Can't create user. Err msg:%v.", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("User id = %d", u.UserID)})
		w.WriteHeader(http.StatusCreated)
	}
}
