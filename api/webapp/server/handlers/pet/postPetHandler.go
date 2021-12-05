package pet

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PostPetHandler creates Pet
func PostPetHandler() http.HandlerFunc {
	pets := dto.GetPetsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &petRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		p := entity.Pet{
			PetID:     req.PetID,
			Name:      req.Name,
			Type:      entity.PetType(req.Type),
			OwnerID:   req.OwnerID,
			Weight:    req.Weight,
			Diesieses: req.Diesieses,
		}

		pets = append(pets, dto.PetDto(entity.PetToDto(p)))
		json.NewEncoder(w).Encode(pets)
	}
}
