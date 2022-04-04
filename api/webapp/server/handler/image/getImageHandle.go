package image

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetImageHandle ...
func GetImageHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "image/jpeg")
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

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while openong DB. Err msg: %v", err)})
			return
		}

		imageDTO, err := s.Image().FindByID(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while searching image. Err msg: %v", err)})
			return
		}

		imageDTO.Format = format

		image, err := s.ImageRepository.GetImageFromLocalStore(imageDTO)

		w.WriteHeader(http.StatusOK)
		jpeg.Encode(w, *image, nil)
	}
}
