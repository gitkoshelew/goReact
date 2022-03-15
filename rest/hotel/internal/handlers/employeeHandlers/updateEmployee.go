package employeehandlers

import (
	"encoding/json"
	"fmt"
	"hotel/domain/model"
	"hotel/internal/apperror"
	"hotel/internal/store"
	"hotel/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UpdateEmployee ...
func UpdateEmployee(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		req := &model.EmployeeDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		hotel, err := s.Hotel().FindByID(req.HotelID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting hotel by id.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Error occured while getting hotel by id. Err msg:%v.", err)))
			return
		}

		employeeDTO, err := s.Employee().FindByID(req.EmployeeID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting employee by id.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Error occured while getting employee by id. Err msg:%v.", err)))

			return
		}

		employee, err := s.Employee().EmployeeFromDTO(employeeDTO)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while converting employeeDTO.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Error occured while converting employeeDTO. Err msg:%v.", err)))
			return
		}

		if hotel != nil {
			if employee.Hotel.HotelID != req.HotelID {
				employee.Hotel = *hotel
			}
		}

		if req.UserID != 0 {
			employee.UserID = req.UserID
		}
		if req.Position != "" {
			employee.Position = req.Position
		}

		err = employee.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Data is not valid.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Data is not valid. Err msg:%v.", err)))
			return
		}

		err = s.Employee().Update(employee)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while updating employee.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Error occured while updating employee. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Updated employee with id = %d", employee.EmployeeID)})

	}
}
