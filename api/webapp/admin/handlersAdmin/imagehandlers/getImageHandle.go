package imagehandlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/handler/response"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetImageHandle ...
func GetImageHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		imgs := []model.ImageDTO{}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while parsing id. Err msg: %v", err)})
			return
		}

		format := r.FormValue("Format")

		err = s.Open()
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			return
		}

		imageDTO, err := s.Image().FindByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting image by id. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}

		imageDTO.Format = format

		image, err := s.ImageRepository.GetImageFromLocalStore(imageDTO)

		imgs = append(imgs, *imageDTO)

		/*files := []string{
			"/api/webapp/admin/tamplates/allImages.html",
			"/api/webapp/admin/tamplates/base.html",
		}*/

		files := []string{
			"/api/webapp/admin/tamplates/image.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while parsing template: %v", err)
			return
		}

		err = tmpl.Execute(w, image)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while executing template: %v", err)
			return
		}
	}
}
