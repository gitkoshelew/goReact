package account

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteAccountHandle deletes account
func DeleteAccountHandle() httprouter.Handle {
	accountsDto := dto.GetAccountsDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	}
}
