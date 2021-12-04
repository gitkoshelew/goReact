package handlers

import (
	"encoding/json"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type petRequest struct {
	PetID     int     `json:"petId"`
	Name      string  `json:"name"`
	Type      string  `json:"petType"`
	OwnerID   int     `json:"ownerId"`
	Weight    float32 `json:"weight"`
	Diesieses string  `json:"diesieses"`
}

// HandlePets  GET /api/pets - returns all pets(JSON)
//	   	  	   POST /api/pet - add pet(JSON)
//			   PUT /api/pet - update pet(JSON)
func HandlePets() http.HandlerFunc {

	pets := entity.GetPetsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(pets)
		// POST
		case http.MethodPost:
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
			pets = append(pets, entity.PetToDto(p))
			json.NewEncoder(w).Encode(pets)
		// PUT
		case http.MethodPut:
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
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandlePet GET /api/pet/:id - returns pet by ID (JSON)
// 		 	 DELETE /api/pet/:id - delete pet by ID(JSON)
func HandlePet() httprouter.Handle {

	pets := entity.GetPetsDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			json.NewEncoder(w).Encode(entity.PetToDto(entity.GetPetByID(id)))
		// DELETE
		case http.MethodDelete:
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
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
