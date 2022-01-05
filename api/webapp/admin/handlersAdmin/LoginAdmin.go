package handlersadmin

import (
	"fmt"
	"html/template"
	"net/http"

	"goReact/webapp/admin/session"

	"github.com/julienschmidt/httprouter"
)

func LoginAdmin() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		exist := session.IsExist(w, r)
		if exist {
			fmt.Println("Exist", exist)
			http.Redirect(w, r, "/admin/home", http.StatusFound)
			return
		}

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
