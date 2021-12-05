package account

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PostAccountHandler creates Account
func PostAccountHandler() http.HandlerFunc {
	accountsDto := dto.GetAccountsDto()

	return func(w http.ResponseWriter, r *http.Request) {

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

		a := dto.AccountDto{
			AccountID: req.AccountID,
			Login:     req.Login,
			Password:  req.Password,
		}
		accountsDto = append(accountsDto, a)
		json.NewEncoder(w).Encode(accountsDto)
		w.WriteHeader(http.StatusCreated)
	}
}
