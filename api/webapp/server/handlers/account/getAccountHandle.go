package account

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetAccountHandle returns account by ID
func GetAccountHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		row := db.QueryRow("SELECT * FROM ACCOUNT WHERE id = $1", id)

		account := dto.AccountDto{}
		err = row.Scan(
			&account.AccountID,
			&account.Login,
			&account.Password)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(account)
	}
}
