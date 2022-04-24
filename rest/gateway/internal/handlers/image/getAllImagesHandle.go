package image

import (
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/image"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAllImagesHandle ...
func GetAllImagesHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		getAllService, err := image.GetAll(r.Context(), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var image []*ImageDTO
		if err := json.Unmarshal(getAllService.Body, &image); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(image)
	}
}
