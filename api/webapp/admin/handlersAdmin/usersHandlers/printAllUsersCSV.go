package usershandlers

import (
	"context"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/admin/middlewear"
	"goReact/webapp/admin/pkg/csv"
	"goReact/webapp/admin/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PrintAllUsersCSV in csv file
func PrintAllUsersCSV(s *store.Store, next http.Handler) httprouter.Handle {
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

		users, err := s.User().GetAll()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting all users. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}
		name := "allusers.csv"

		path, err := csv.MakeCSV(users, name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while recording to csv. Err msg:%v. ", err), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while recording to csv. Err msg:%v.", err)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), middlewear.CtxKeyFile, path)))
	}
}
