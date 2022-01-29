package room

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/handler/pagination"
	"goReact/webapp/server/handler/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetRoomsHandlePagination returns all rooms with limit and offset
func GetRoomsHandlePagination(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		page := &pagination.Page{}

		var err error
		page.PageNumber, err = strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad Query. Err msg:%v: ", err)

		}

		page.PageSize, err = strconv.Atoi(r.URL.Query().Get("pagesize"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad Query. Err msg:%v: ", err)

		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})

		}

		rooms, err := s.Room().GetAllPagination(page)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Can't find user. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rooms)
	}
}
