package image

import (
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/image"
	"gateway/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UpdateimageHandle ...
func UpdateimageHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		updateService, err := image.Update(r.Context(), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var response *response.Info
		if err := json.Unmarshal(updateService.Body, &response); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
