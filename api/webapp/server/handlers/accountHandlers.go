package handlers

import (
	"encoding/json"
	"goReact/domain/dto"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type accountRequest struct {
	AccountID int    `json:"accountId"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

// HandleAccounts  GET /api/accounts - returns all accounts(JSON)
//				   POST /api/account - add account(JSON)
//				   PUT /api/account - update account(JSON)
func HandleAccounts() http.HandlerFunc {

	accountsDto := entity.GetAccountsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(accountsDto)
		//POST
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &accountRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}
			if len(req.Login) < 2 || len(req.Password) < 2 {
				http.Error(w, "Login and password should have at least 3 symbols", http.StatusBadRequest)
				return
			}
			for _, v := range accountsDto {
				if req.AccountID == v.AccountID || req.Login == v.Login {
					http.Error(w, "Login or ID is already taken, try another", http.StatusBadRequest)
					return
				}
			}

			a := dto.Account{
				AccountID: req.AccountID,
				Login:     req.Login,
				Password:  req.Password,
			}
			accountsDto = append(accountsDto, a)
			json.NewEncoder(w).Encode(accountsDto)
			w.WriteHeader(http.StatusCreated)
		// PUT
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &accountRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}
			if len(req.Login) < 2 || len(req.Password) < 2 {
				http.Error(w, "Password should have at least 3 symbols", http.StatusBadRequest)
				return
			}

			for index, acc := range accountsDto {
				if acc.AccountID == req.AccountID {
					if accountsDto[index].Password == req.Password {
						http.Error(w, "New password cannot can't do match the old password", http.StatusBadRequest)
						return
					}
					accountsDto[index].Password = req.Password
					break
				}
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(accountsDto)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleAccount GET /api/account/:id - returns account by ID (JSON)
// 				 DELETE /api/account/:id - delete account by ID(JSON)
func HandleAccount() httprouter.Handle {

	accountsDto := entity.GetAccountsDto()

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

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(entity.AccountToDto(entity.GetAccountByID(id)))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, acc := range accountsDto {
				if acc.AccountID == id { // delete object imitation =)
					accountsDto[index].AccountID = 0
					accountsDto[index].Login = "NIL"
					accountsDto[index].Password = "NIL"
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(accountsDto)
					return
				}
			}

			http.Error(w, "Cant find account", http.StatusBadRequest)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
