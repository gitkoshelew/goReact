package usershandlers

import (
	"customer/domain/model"
	"customer/internal/apperror"
	"customer/internal/store"
	"customer/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

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
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		u, err := s.User().FindByID(req.UserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting user by id", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
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
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Data is not valid.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Data is not valid. Err msg:%v.", err)))
			return
		}

		err = s.User().Update(u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while updating user.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Error occured while updating user. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Updated user with id = %d", u.UserID)})

	}
}
