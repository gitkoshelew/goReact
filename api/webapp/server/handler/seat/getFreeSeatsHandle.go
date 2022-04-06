package seat

import (
	"encoding/json"
	"fmt"
	reqandresp "goReact/domain/reqAndResp"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"net/http"
)

// GetFreeSeatsHandle ...
func GetFreeSeatsHandle(s *store.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		request := r.Context().Value(handler.CtxKeyFreeSeatsSearchReqValidation).(*reqandresp.FreeSeatsSearching)

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while opening DB. Err msg: %v", err)})
			return
		}

		seatIds, err := s.Seat().FreeSeatsSearching(request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while searching seat. Err msg: %v", err)})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(seatIds)
	})
}
