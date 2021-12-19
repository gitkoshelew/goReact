package image

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetImageHandle returns image by ID
func GetImageHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &imageRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		log.Printf("%d + %s", req.ImageID, req.Type)

		row := db.QueryRow("SELECT * FROM IMAGES WHERE id = $1 and type = $2", req.ImageID, req.Type)

		image := dto.ImageDto{}
		err := row.Scan(
			&image.ImageID,
			&image.Type,
			&image.URL,
			&image.OwnerID)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(image)
	}
}
