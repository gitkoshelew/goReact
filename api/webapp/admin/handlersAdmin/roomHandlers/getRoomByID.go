package roomhandlers

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetRoomByID ...
func GetRoomByID(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		rooms := []model.RoomDTO{}
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}

		room, err := s.Room().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find pet. Err msg:%v.", err)
			return
		}
		rooms = append(rooms, *room)

		files := []string{
			"/api/webapp/admin/tamplates/allRooms.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}

		err = tmpl.Execute(w, rooms)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}
	}
}
