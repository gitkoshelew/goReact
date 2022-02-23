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

// UpdateUser ...
func UpdateUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		req := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}

		u, err := s.User().FindByID(req.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find user. Err msg:%v.", err)
			return
		}

		if req.Email != "" {
			u.Email = req.Email
		}

		if req.Role != "" {
			u.Role = req.Role
		}

		if req.Name != "" {
			u.Name = req.Name
		}

		if req.Surname != "" {
			u.Surname = req.Surname
		}

		if req.MiddleName != "" {
			u.MiddleName = req.MiddleName
		}

		if req.Sex != "" {
			u.Sex = req.Sex
		}

		if !req.DateOfBirth.IsZero() {
			u.DateOfBirth = req.DateOfBirth
		}

		if req.Address != "" {
			u.Address = req.Address
		}

		if req.Phone != "" {
			u.Phone = req.Phone
		}

		if req.Photo != "" {
			u.Photo = req.Photo
		}

		err = u.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.User().Update(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't update user. Err msg:%v.", err)
			return
		}

		s.Logger.Info("Update user with id = %d", u.UserID)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Update user with id = %d", u.UserID)})

	}
}
