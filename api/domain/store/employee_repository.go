package store

import (
	"goReact/domain/model"
	"log"
)

// EmployeeRepository ...
type EmployeeRepository struct {
	Store *Store
}

// Create employee and save it to DB
func (r *EmployeeRepository) Create(e *model.Employee) (*model.Employee, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO employee",
		"(user_id, hotel_id, position, role)",
		"VALUES ($1, $2, $3, $4) RETURNING id",
		e.UserID,
		e.Hotel.HotelID,
		e.Position,
	).Scan(&e.EmployeeID); err != nil {
		log.Print(err)
		return nil, err
	}
	return e, nil
}

// GetAll returns all employees
func (r *EmployeeRepository) GetAll() (*[]model.Employee, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM employee")
	if err != nil {
		log.Print(err)
	}
	employees := []model.Employee{}

	for rows.Next() {
		employee := model.Employee{}
		err := rows.Scan(
			&employee.EmployeeID,
			&employee.UserID,
			&employee.Hotel.HotelID,
			&employee.Position,
		)
		if err != nil {
			log.Printf(err.Error())
			continue
		}
		employees = append(employees, employee)
	}
	return &employees, nil
}

// FindByID searchs and returns employee by ID
func (r *EmployeeRepository) FindByID(id int) (*model.Employee, error) {
	employee := &model.Employee{}
	if err := r.Store.Db.QueryRow("SELECT * FROM employee WHERE id = $1",
		id).Scan(
		&employee.EmployeeID,
		&employee.UserID,
		&employee.Hotel.HotelID,
		&employee.Position,
	); err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return employee, nil
}

// Delete employee from DB by ID
func (r *EmployeeRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM employee WHERE id = $1", id)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	log.Printf("Employee deleted, rows affectet: %d", result)
	return nil
}

// Update employee from DB
func (r *EmployeeRepository) Update(e *model.Employee) error {

	result, err := r.Store.Db.Exec(
		"UPDATE employee SET",
		"user_id = $1, hotel_id = $2, position = $3",
		"WHERE id = $4",
		e.UserID,
		e.Hotel.HotelID,
		e.Position,
		e.EmployeeID,
	)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	log.Printf("Employee updated, rows affectet: %d", result)
	return nil
}
