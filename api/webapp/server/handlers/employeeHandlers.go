package handlers

import (
	"encoding/json"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type employeeRequest struct {
	UserID     int    `json:"userId"`
	HotelID    int    `json:"hotelId"`
	EmployeeID int    `json:"employeeId"`
	Position   string `json:"position"`
	Role       string `json:"role"`
}

// HandleEmployees  GET /api/employees - returns all employees(JSON)
//			  	    POST /api/employee - add employee(JSON)
//			 	    PUT /api/employee - update employee(JSON)
func HandleEmployees() http.HandlerFunc {

	employeesDto := entity.GetEmployeesDto()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(employeesDto)
		// POST
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &employeeRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			e := entity.Employee{
				User:       entity.GetUserByID(req.UserID),
				Hotel:      entity.GetHotelByID(req.HotelID),
				EmployeeID: req.EmployeeID,
				Position:   req.Position,
				Role:       req.Role,
			}
			employeesDto = append(employeesDto, entity.EmployeeToDto(e))
			json.NewEncoder(w).Encode(employeesDto)
		// PUT
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &employeeRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, e := range employeesDto {
				if e.EmployeeID == req.EmployeeID {
					employeesDto[index].HotelID = req.HotelID
					employeesDto[index].Position = req.Position
					employeesDto[index].Role = req.Role
					break
				}
			}
			json.NewEncoder(w).Encode(employeesDto)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleEmployee GET /api/employee/:id - returns employee by ID (JSON)
// 				  DELETE /api/employee/:id - delete employee by ID(JSON)
func HandleEmployee() httprouter.Handle {

	employeesDto := entity.GetEmployeesDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			json.NewEncoder(w).Encode(entity.EmployeeToDto(entity.GetEmployeeByID(id)))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, e := range employeesDto {
				if e.EmployeeID == id { // delete object imitation =)
					employeesDto[index].Position = "DELETE"
					employeesDto[index].Role = "DELETE"
					json.NewEncoder(w).Encode(employeesDto)
					return
				}
			}
			http.Error(w, "Cant find Employee", http.StatusBadRequest)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
