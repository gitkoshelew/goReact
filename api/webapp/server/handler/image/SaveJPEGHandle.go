package image

import (
	"encoding/json"
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// imageInfo ...
type imageInfo struct {
	Name string `json:"fileName"`
}

// SaveJPEGHandle ...
func SaveJPEGHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		mr, err := r.MultipartReader()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("error occured while reading multi form data request. Err msg: %v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error occured while reading multi form data request. Err msg: %v", err)})
			return
		}

		imageDTO := &model.ImageDTO{}
		var image image.Image

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
					json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error occured while decoding image. Err msg: %v.", err)})
					return
				}
			} else if part.FormName() == "type" {
				buffer := make([]byte, 32)
				n, _ := part.Read(buffer)
				if n == 0 {
					w.WriteHeader(http.StatusBadRequest)
					s.Logger.Error("error occured while decoding type. Empty value of key 'type'")
					json.NewEncoder(w).Encode(response.Error{Messsage: "error occured while decoding type. Empty value of key 'type'"})
					return
				}
				imageDTO.Type = string(buffer[:n])

			} else if part.FormName() == "ownerId" {
				buffer := make([]byte, 32)
				n, _ := part.Read(buffer)
				if n == 0 {
					w.WriteHeader(http.StatusBadRequest)
					s.Logger.Error("error occured while decoding type. Empty value of key 'ownerId'")
					json.NewEncoder(w).Encode(response.Error{Messsage: "error occured while decoding type. Empty value of key 'ownerId'"})
					return
				}
				id, err := strconv.Atoi(string(buffer[:n]))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					s.Logger.Errorf("error occured while convertig string to int. Err msg: %v.", err)
					json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error occured while convertig string to int. Err msg: %v.", err)})
					return
				}
				imageDTO.OwnerID = id
			}

		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while openong DB. Err msg: %v", err)})
			return
		}

		_, imageURL, err := s.Image().SaveImage(imageDTO, &image)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("eror occured while saving JPEG.  Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		URLs := model.BuildImagesURLsResponse(imageURL)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(URLs)
	}
}
