package account

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PutAccountHandler updates Account
func PutAccountHandler() http.HandlerFunc {
	accountsDto := dto.GetAccountsDto()

	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}
