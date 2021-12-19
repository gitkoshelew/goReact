package account

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAccountsHandle returns all accounts
func GetAccountsHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT * FROM ACCOUNT")
		if err != nil {
			log.Fatal(err)
		}

		accounts := []dto.AccountDto{}

		for rows.Next() {
			account := dto.AccountDto{}
			err := rows.Scan(
				&account.AccountID,
				&account.Login,
				&account.Password)

			if err != nil {
				log.Printf(err.Error())
				continue
			}
			accounts = append(accounts, account)
		}

		json.NewEncoder(w).Encode(accounts)
	}
}
