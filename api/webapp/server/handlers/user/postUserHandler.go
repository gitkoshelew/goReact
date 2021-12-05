package user

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PostUserHandler creates User
func PostUserHandler() http.HandlerFunc {
	users := dto.GetUsersDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &userRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		for _, v := range users {
			if req.Email == v.Email {
				http.Error(w, "Email is already taken, try another", http.StatusBadRequest)
				return
			}
		}

		u := entity.User{
			Account:     entity.GetAccountByID(req.AccountID),
			UserID:      req.UserID,
			Name:        req.Name,
			Surname:     req.Surname,
			MiddleName:  req.MiddleName,
			DateOfBirth: req.DateOfBirth,
			Address:     req.Address,
			Phone:       req.Phone,
			Email:       req.Email,
		}
		users = append(users, dto.UserDto(entity.UserToDto(u)))
		json.NewEncoder(w).Encode(users)
		w.WriteHeader(http.StatusCreated)
	}
}
