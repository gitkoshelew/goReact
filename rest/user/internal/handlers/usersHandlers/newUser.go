package usershandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/domain/model"
	"user/internal/store"
	"user/pkg/response"

	"github.com/julienschmidt/httprouter"
)

// NewUser ...
func NewUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
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
			Verified:    false,
			DateOfBirth: req.DateOfBirth,
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}

		err = u.WithEncryptedPassword()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v.", err)
			return
		}

		err = u.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		_, err = s.User().Create(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't create user. Err msg:%v.", err)
			return
		}

		s.Logger.Info("Creat user with id = %d", u.UserID)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Creat user with id = %d", u.UserID)})
	}
}
