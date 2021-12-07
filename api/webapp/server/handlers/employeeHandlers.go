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

	employees := entity.GetEmployees()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(employees)
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
			employees = append(employees, e)
			json.NewEncoder(w).Encode(employees)
		// PUT
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &employeeRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, e := range employees {
				if e.EmployeeID == req.EmployeeID {
					employees[index].User = entity.GetUserByID(req.UserID)
					employees[index].Hotel = entity.GetHotelByID(req.HotelID)
					employees[index].EmployeeID = req.EmployeeID
					employees[index].Position = req.Position
					employees[index].Role = req.Role
					break
				}
			}
			json.NewEncoder(w).Encode(employees)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleEmployee GET /api/employee/:id - returns employee by ID (JSON)
// 				  DELETE /api/employee/:id - delete employee by ID(JSON)
func HandleEmployee() httprouter.Handle {

	employees := entity.GetEmployees()

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

			json.NewEncoder(w).Encode(entity.GetEmployeeByID(id))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, e := range employees {
				if e.EmployeeID == id { // delete object imitation =)
					employees[index].Position = "DELETE"
					employees[index].Role = "DELETE"
					json.NewEncoder(w).Encode(employees)
					return
				}
			}
			http.Error(w, "Cant find Employee", http.StatusBadRequest)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
