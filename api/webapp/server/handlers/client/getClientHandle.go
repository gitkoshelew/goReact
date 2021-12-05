package client

import (
	"encoding/json"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetClientHandle returns Client by ID
func GetClientHandle() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		id, err := strconv.Atoi(ps.ByName("id"))

		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		json.NewEncoder(w).Encode(entity.ClientToDto(entity.GetClientByID(id)))
	}
}
