package imagehandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var permissionRead model.Permission = model.Permission{Name: model.ReadImage}

// GetAllImagesHandle ...
func GetAllImagesHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionRead.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		imgs, err := s.Image().GetAll()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting all images. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/allImages.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while parsing template: %v", err)
			return
		}

		err = tmpl.Execute(w, imgs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while executing template: %v", err)
			return
		}
	}
}
