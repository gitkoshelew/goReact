package imagehandlers

import (
	"admin/domain/store"
	"admin/webapp/middlewear"
	"admin/webapp/session"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DownloandImage
func DownloandImage(s *store.Store, next http.Handler) httprouter.Handle {
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
		path := fmt.Sprintf("images/%s/id-%d/image-%s.jpg", imageDTO.Type, imageDTO.OwnerID, imageDTO.Format)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), middlewear.CtxKeyFile, path)))
	}
}
