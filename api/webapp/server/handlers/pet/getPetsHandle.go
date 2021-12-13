package pet

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetPetsHandle returns all Pets
func GetPetsHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT * FROM PET")
		if err != nil {
			log.Fatal(err)
		}

		petsDto := []dto.PetDto{}

		for rows.Next() {
			pet := dto.PetDto{}
			err := rows.Scan(
				&pet.PetID,
				&pet.Name,
				&pet.Type,
				&pet.Weight,
				&pet.Diesieses,
				&pet.OwnerID)

			if err != nil {
				log.Printf(err.Error())
				continue
			}

			petsDto = append(petsDto, pet)
		}

		json.NewEncoder(w).Encode(petsDto)
	}
}
