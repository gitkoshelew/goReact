package account

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PostAccountHandle creates Account
func PostAccountHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &accountRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		encryptedPassword, err := store.EncryptPassword(req.Password)
		if err != nil {
			log.Fatal(err)
		}

		var id int
		err = db.QueryRow("INSERT into ACCOUNT (Login, Password) VALUES ($1, $2) RETURNING id",
			req.Login, encryptedPassword,
		).Scan(&id)

		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(id)
		w.WriteHeader(http.StatusCreated)
	}
}
