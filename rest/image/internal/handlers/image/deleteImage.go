package image

import (
	"encoding/json"
	"fmt"
	"image/domain/store"
	"image/pkg/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteImageHandle ...
func DeleteImageHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while parsing id. Err msg: %v", err)})
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while openong DB. Err msg: %v", err)})
			return
		}

		err = s.Image().Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while deleting image. Err msg: %v", err)})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("image with id %d deleted", id)})

	}
}
