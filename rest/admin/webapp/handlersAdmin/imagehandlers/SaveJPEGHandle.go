package imagehandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permissionCreate model.Permission = model.Permission{Name: model.CreatImage}

// SaveJPEGHandle ...
func SaveJPEGHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionCreate.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		r.ParseMultipartForm(32000000)

		file, _, err := r.FormFile("Photo")
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.Body), http.StatusInternalServerError)
			s.Logger.Errorf("error occured while reading multi form data request. Err msg: %v. Requests body: %v", err, r.Body)

		}

		id, err := strconv.Atoi(r.FormValue("OwnerID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("OwnerID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("OwnerID"))
			return
		}

		imgtype := r.FormValue("Type")

		imageDTO := &model.ImageDTO{
			OwnerID: id,
			Type:    imgtype,
		}
		var image image.Image

		image, err = jpeg.Decode(file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while decoding file to jpeg. Err msg: %v.", err), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while decoding file to jpeg. Err msg: %v.", err)
			return
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = imageDTO.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		_, err = s.Image().SaveImage(imageDTO, &image)
		if err != nil {
			http.Error(w, fmt.Sprintf("eror occured while saving JPEG.  Err msg: %v", err), http.StatusInternalServerError)
			s.Logger.Errorf("eror occured while saving JPEG.  Err msg: %v", err)
			return
		}
		http.Redirect(w, r, "/admin/homeimages/", http.StatusFound)

	}

}
