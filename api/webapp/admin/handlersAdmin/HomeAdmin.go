package handlersadmin

import (
	"goReact/webapp/admin/session"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HomeAdmin ...
func HomeAdmin() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		exist := session.IsExist(w, r)
		if exist {
			HomePage(w)
			return
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
		}
	}
}

func HomePage(w http.ResponseWriter) {
	files := []string{
		"/api/webapp/admin/tamplates/homeAdmin.html",
		"/api/webapp/admin/tamplates/base.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
