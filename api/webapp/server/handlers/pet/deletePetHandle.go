package pet

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeletePetHandle deletes Pet
func DeletePetHandle() httprouter.Handle {
	pets := dto.GetPetsDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, p := range pets {
			if p.PetID == id { // delete object imitation =)
				pets[index].Name = "DELETE"
				json.NewEncoder(w).Encode(pets)
				return
			}
		}

		http.Error(w, "Cant find Pet", http.StatusBadRequest)
	}
}
