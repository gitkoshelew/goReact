package imagehandlers

import (
	"fmt"
	"image/jpeg"

	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetImageHandle ...
func GetImageHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionRead.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			return
		}

		format := r.FormValue("Format")

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		imageDTO, err := s.Image().FindByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting image by id. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}

		imageDTO.Format = format

		image, err := s.ImageRepository.GetImageFromLocalStore(imageDTO)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting image file. Err msg: %v.", err), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while getting image file. Err msg: %v:", err)
			return
		}

		jpeg.Encode(w, *image, nil)
	}
}
