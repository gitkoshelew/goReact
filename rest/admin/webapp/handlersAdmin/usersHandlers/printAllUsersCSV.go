package usershandlers

import (
	"admin/domain/store"
	"admin/pkg/csv"
	"admin/webapp/middlewear"
	"admin/webapp/session"
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PrintAllUsersCSV in csv file
func PrintAllUsersCSV(s *store.Store, next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_read.Name)
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

		users, err := s.User().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		name := "allusers.csv"

		path, err := csv.MakeCSV(users, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Error occured while recording to csv:", err)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), middlewear.CtxKeyFile, path)))
	}
}
