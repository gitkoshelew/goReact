package account

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// GetAccountsHandler returns all accounts
func GetAccountsHandler() http.HandlerFunc {
	accountsDto := dto.GetAccountsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accountsDto)
	}
}
