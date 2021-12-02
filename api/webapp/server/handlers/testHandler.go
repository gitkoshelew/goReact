package handlers

import (
	"encoding/json"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Test ...
func Test() httprouter.Handle {
	accounts := entity.GetAccounts()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		id, err := strconv.Atoi(ps.ByName("id"))

		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for _, acc := range accounts {
			if acc.AccountID == id {
				json.NewEncoder(w).Encode(acc)
			}
		}
	}
}
