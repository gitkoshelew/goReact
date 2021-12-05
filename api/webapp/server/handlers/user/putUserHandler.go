package user

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PutUserHandle updates User
func PutUserHandle() http.HandlerFunc {

	users := dto.GetUsersDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &userRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, user := range users {
			if user.UserID == req.UserID {
				users[index].Name = req.Name
				users[index].Surname = req.Surname
				users[index].MiddleName = req.MiddleName
				users[index].DateOfBirth = req.DateOfBirth
				users[index].Address = req.Address
				users[index].Phone = req.Phone
				break
			}
		}

		json.NewEncoder(w).Encode(users)
		w.WriteHeader(http.StatusCreated)
	}
}
