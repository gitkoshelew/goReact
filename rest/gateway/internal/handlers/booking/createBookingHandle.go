package booking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/internal/apperror"
	"gateway/internal/client"
	"gateway/internal/client/booking"
	"gateway/internal/client/utils"
	"gateway/pkg/response"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateBookingHandle ...
func CreateBookingHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		rBody, _ := ioutil.ReadAll(r.Body)

		req := &DataValidation{}
		if err := json.NewDecoder(bytes.NewReader(rBody)).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			service.Base.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", bytes.NewReader(rBody), err)
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("Eror during JSON request decoding. Request body: %v", bytes.NewReader(rBody)), fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		var dataValidation = map[utils.DataValidation]int{
			utils.PetValidation:      req.PetID,
			utils.EmployeeValidation: req.EmployeeID,
			utils.SeatValidation:     req.SeatID,
		}

		err := utils.IsValid(bytes.NewReader(rBody), dataValidation, service.Base.Logger.Logger)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		createService, err := booking.Create(r.Context(), service, bytes.NewReader(rBody))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var response *response.Info
		if err := json.Unmarshal(createService.Body, &response); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
