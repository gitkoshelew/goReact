package imagehandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"fmt"

	"admin/webapp/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permissionDelete model.Permission = model.Permission{Name: model.DeleteImage}

// DeleteImageHandle ...
func DeleteImageHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionDelete.Name)
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

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = s.Image().Delete(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while deleting hotel. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/homeimages", http.StatusFound)

	}
}
