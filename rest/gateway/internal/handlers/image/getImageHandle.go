package image

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/internal/client"
	"gateway/internal/client/image"
	"gateway/pkg/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetImageHandle ...
func GetImageHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while parsing id. Err msg: %v", err)})
			return
		}
		format := r.URL.Query().Get("format")
		if format == "" {
			format = "original"
		}
		getService, err := image.Get(context.WithValue(r.Context(), client.ImageGetQuerryParamsCtxKey, fmt.Sprintf("?id=%d&format=%s", id, format)), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var image *ImageDTO
		if err := json.Unmarshal(getService.Body, &image); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(image)
	}
}
