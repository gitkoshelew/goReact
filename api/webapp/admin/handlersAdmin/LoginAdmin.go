package handlersadmin

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LoginAdmin ...
func LoginAdmin() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		files := []string{
			"/api/webapp/admin/tamplates/loginAdmin.html",
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
}
