package client

import (
	"goReact/webapp/server/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteClientHandle delete Client by ID
func DeleteClientHandle() httprouter.Handle {

	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		result, err := db.Exec("DELETE from CLIENT WHERE id = $1", id)

		if err != nil {
			panic(err)
		}

		log.Print(result.RowsAffected())
	}
}
