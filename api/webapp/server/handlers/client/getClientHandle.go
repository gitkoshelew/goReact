package client

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetClientHandle returns Client by ID
func GetClientHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		id, err := strconv.Atoi(ps.ByName("id"))

		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		row := db.QueryRow("SELECT * FROM CLIENT WHERE id = $1", id)

		client := dto.ClientDto{}
		err = row.Scan(
			&client.ClientID,
			&client.UserID)

		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(client)
	}
}
