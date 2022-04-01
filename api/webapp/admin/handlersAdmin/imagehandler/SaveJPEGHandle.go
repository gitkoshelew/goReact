package imagehandler

import (
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// SaveJPEGHandle ...
func SaveJPEGHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		r.ParseMultipartForm(32 << 20)
		//https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.5.html

		file, header, err := r.FormFile("Photo")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.Body), http.StatusInternalServerError)
			s.Logger.Errorf("error occured while reading multi form data request. Err msg: %v. Requests body: %v", err, r.Body)

		}

		s.Logger.Info("Header: %v ", header)

		/*mr, err := r.MultipartReader()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("error occured while reading multi form data request. Err msg: %v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error occured while reading multi form data request. Err msg: %v", err)})
		}*/
		/*
			for {
				part, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					s.Logger.Errorf("error occured while reading parts of multipart reader. Err msg: %v", err)
					json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error occured while reading parts of multipart reader. Err msg: %v", err)})
				}

				if part.FormName() == "image" {
					image, err = jpeg.Decode(part)

					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						s.Logger.Errorf("error occured while decoding image. Err msg: %v.", err)
						return
					}
				} else if part.FormName() == "json" {
					if err := json.NewDecoder(part).Decode(imageDTO); err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						s.Logger.Errorf("eror occured while JSON request decoding. Request body: %v, Err msg: %v", part, err)
						json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
						return
					}
				}
			}
		*/

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

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = s.Image().SaveImage(imageDTO, &image)
		if err != nil {
			http.Error(w, fmt.Sprintf("eror occured while saving JPEG.  Err msg: %v", err), http.StatusInternalServerError)
			s.Logger.Errorf("eror occured while saving JPEG.  Err msg: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		http.Redirect(w, r, "/admin/imagehome/", http.StatusFound)

	}

}
