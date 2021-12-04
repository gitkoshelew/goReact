package handlers

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/pkg/date"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type userRequest struct {
	AccountID   int       `json:"accountId"`
	UserID      int       `json:"userId"`
	Name        string    `json:"name"`
	Surname     string    `json:"sName"`
	MiddleName  string    `json:"mName"`
	DateOfBirth date.Date `json:"birthDate"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
}

// HandleUsers  GET /api/users - returns all users(JSON)
//			    POST /api/user - add user(JSON)
//			    PUT /api/user - update user(JSON)
func HandleUsers() http.HandlerFunc {

	users := entity.GetUsersDto()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users)
		// POST
		case http.MethodPost:
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
			users = append(users, entity.UserToDto(u))
			json.NewEncoder(w).Encode(users)
			w.WriteHeader(http.StatusCreated)
		// PUT
		case http.MethodPut:
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
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleUser GET /api/user/:id - returns user by ID (JSON)
// 				 DELETE /api/user/:id - delete user by ID(JSON)
func HandleUser() httprouter.Handle {

	users := entity.GetUsersDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			json.NewEncoder(w).Encode(entity.UserToDto(entity.GetUserByID(id)))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, u := range users {
				if u.UserID == id { // delete object imitation =)
					users[index].Name = "DELETE"
					users[index].Surname = "DELETE"
					users[index].MiddleName = "DELETE"
					json.NewEncoder(w).Encode(users)
					return
				}
			}
			http.Error(w, "Cant find User", http.StatusBadRequest)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
