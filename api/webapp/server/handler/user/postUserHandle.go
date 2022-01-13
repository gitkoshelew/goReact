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
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			w.WriteHeader(http.StatusBadRequest)
		}

		u := model.User{
			UserID:      0,
			Email:       req.Email,
			Password:    req.Password,
			Role:        model.Role(req.Role),
			Name:        req.Name,
			Surname:     req.Surname,
			MiddleName:  req.MiddleName,
			Sex:         model.Sex(req.Sex),
			Address:     req.Address,
			Phone:       req.Phone,
			Photo:       req.Photo,
			Verified:    req.Verified,
			DateOfBirth: req.DateOfBirth,
		}
		err := u.Validate()
		if err != nil {
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			w.WriteHeader(http.StatusBadRequest)
		}

		err = u.NewUser()
		if err != nil {
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			w.WriteHeader(http.StatusBadRequest)
		}

		err = s.Open()
		if err != nil {
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = s.User().Create(&u)
		if err != nil {
			s.Logger.Errorf("Can't create user. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("User id = %d", u.UserID)})
		w.WriteHeader(http.StatusCreated)
	}
}
