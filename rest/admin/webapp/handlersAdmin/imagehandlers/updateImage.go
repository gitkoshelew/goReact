package imagehandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permissionUpdate model.Permission = model.Permission{Name: model.UpdateImage}

// UpdateImage ...
func UpdateImage(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionUpdate.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		id, err := strconv.Atoi(r.FormValue("ImageID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("ImageID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("ImageID"))
			return
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		imageDTO, err := s.Image().FindByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while finding image by id. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}

		ownerID, err := strconv.Atoi(r.FormValue("OwnerID"))
		if err == nil {
			if ownerID != 0 {
				imageDTO.OwnerID = ownerID
			}

		}

		imgType := r.FormValue("Type")
		if imgType != "" {
			imageDTO.Type = imgType

		}

		err = imageDTO.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		image, err := s.Image().ModelFromDTO(imageDTO)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while converting DTO. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		err = s.Image().Update(image)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while updating image. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homeimages", http.StatusFound)

	}
}
