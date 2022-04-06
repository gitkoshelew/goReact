package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	reqandresp "goReact/domain/reqAndResp"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ValidateFreeSeatsSearchingRequest ...
func ValidateFreeSeatsSearchingRequest(next http.Handler, s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		request := &reqandresp.FreeSeatsSearching{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err := request.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while validating user. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while validating user. Err msg: %v", err)})
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(context.Background(), handler.CtxKeyFreeSeatsSearchReqValidation, request)))
	}

}
