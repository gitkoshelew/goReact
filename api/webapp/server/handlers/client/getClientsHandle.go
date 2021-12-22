package client

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetClientsHandle returns all Clients
func GetClientsHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT * FROM CLIENT")
		if err != nil {
			log.Fatal(err)
		}

		clientsDto := []dto.ClientDto{}

		for rows.Next() {
			client := dto.ClientDto{}
			err := rows.Scan(
				&client.ClientID,
				&client.UserID)

			if err != nil {
				log.Printf(err.Error())
				continue
			}

			clientsDto = append(clientsDto, client)
		}

		json.NewEncoder(w).Encode(clientsDto)
	}
}
