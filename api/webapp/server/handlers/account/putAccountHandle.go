package account

import (
	"encoding/json"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PutAccountHandle updates Account
func PutAccountHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &accountRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		result, err := db.Exec("UPDATE ACCOUNT set password = $1 WHERE id = $2",
			req.Password, req.AccountID)

		if err != nil {
			panic(err)
		}

		log.Println(result.RowsAffected())

		w.WriteHeader(http.StatusCreated)
	}
}
