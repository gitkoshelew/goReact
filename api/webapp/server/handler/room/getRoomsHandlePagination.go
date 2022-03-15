package room

import (
	"encoding/json"
	"fmt"
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

		pageNumber, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while parsing offset. Err msg: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while parsing offset. Err msg: %v", err)})
			return
		}

		pageSize, err := strconv.Atoi(r.URL.Query().Get("pagesize"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while parsing pagesize. Err msg: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while parsing pagesize. Err msg: %v", err)})
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while opening DB. Err msg: %v", err)})
			return
		}

		page := &pagination.Page{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		}

		rooms, err := s.Room().GetAllPagination(page)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while searching rooms. Err msg: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while searching rooms. Err msg: %v", err)})
			return
		}

		count, err := s.Room().GetTotalRows()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while calculating rows. Err msg: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while calculating rows. Err msg: %v", err)})
			return
		}

		res := make(map[string]interface{})
		res["rooms"] = rooms
		res["totalCount"] = count

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
