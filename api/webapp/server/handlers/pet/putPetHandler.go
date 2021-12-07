package pet

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PutPetHandler updates Pet
func PutPetHandler() http.HandlerFunc {
	pets := dto.GetPetsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &petRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, p := range pets {
			if p.PetID == req.PetID {
				pets[index].Name = req.Name
				pets[index].Type = req.Type
				pets[index].OwnerID = req.OwnerID
				pets[index].Weight = req.Weight
				pets[index].Diesieses = req.Diesieses
				break
			}
		}

		json.NewEncoder(w).Encode(pets)
	}
}
