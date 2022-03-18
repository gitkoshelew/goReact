package store

import (
	"goReact/domain/model"
)

// EmployeeRepository ...
type EmployeeRepository struct {
	Store *Store
}

// Create employee and save it to DB
func (r *EmployeeRepository) Create(e *model.Employee) (*int, error) {
	if err := r.Store.Db.QueryRow("INSERT INTO employee (user_id, hotel_id, position) VALUES ($1, $2, $3) RETURNING id",
		e.UserID,
		e.Hotel.HotelID,
		e.Position,
	).Scan(&e.EmployeeID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating employee. Err msg:%v.", err)
		return nil, err
	}

	r.Store.Logger.Infof("Employee with id %d was created.", e.EmployeeID)

	return &e.EmployeeID, nil
}

// GetAll returns all employees
func (r *EmployeeRepository) GetAll() (*[]model.EmployeeDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM employee")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all employees. Err msg: %v", err)
	}
	employees := []model.EmployeeDTO{}

	for rows.Next() {
		employee := model.EmployeeDTO{}
		err := rows.Scan(
			&employee.EmployeeID,
			&employee.UserID,
			&employee.HotelID,
			&employee.Position,
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occured while getting all employees. Err msg: %v", err)
			continue
		}
		employees = append(employees, employee)
	}
	return &employees, nil
}

// FindByID searchs and returns employee by ID
func (r *EmployeeRepository) FindByID(id int) (*model.EmployeeDTO, error) {
	employeeDTO := &model.EmployeeDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM employee WHERE id = $1", id).Scan(
		&employeeDTO.EmployeeID,
		&employeeDTO.UserID,
		&employeeDTO.HotelID,
		&employeeDTO.Position,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting employee by id. Err msg: %v.", err)
		return nil, err
	}

	return employeeDTO, nil
}

// Delete employee from DB by ID
func (r *EmployeeRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM employee WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting employee. Err msg: %v.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting employee. Err msg: %v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while deleting employee. Err msg: %v.", err)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Employee with id %d was deleted", id)
	return nil
}

// Update employee from DB
func (r *EmployeeRepository) Update(e *model.EmployeeDTO) error {
	result, err := r.Store.Db.Exec(
		"UPDATE employee SET",
		"user_id = $1, hotel_id = $2, position = $3",
		"WHERE id = $4",
		e.UserID,
		e.HotelID,
		e.Position,
		e.EmployeeID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Erorr occured while updating booking. Err msg: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating employee. Err msg: %v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while updating employee. Err msg: %v.", err)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Employee with id %d was updated", e.EmployeeID)
	return nil
}

//FindByUserID find employee by user ID
func (r *EmployeeRepository) FindByUserID(userID int) (*model.Employee, error) {
	employeeDTO := &model.EmployeeDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM employee WHERE user_id = $1", userID).Scan(
		&employeeDTO.EmployeeID,
		&employeeDTO.UserID,
		&employeeDTO.HotelID,
		&employeeDTO.Position,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while searching employee by user id. Err msg: %v.", err)
		return nil, err
	}

	employee, err := r.ModelFromDTO(employeeDTO)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

// ModelFromDTO ...
func (r *EmployeeRepository) ModelFromDTO(dto *model.EmployeeDTO) (*model.Employee, error) {
	hotel, err := r.Store.Hotel().FindByID(dto.HotelID)
	if err != nil {
		return nil, err
	}

	userDTO, err := r.Store.User().FindByID(dto.UserID)
	if err != nil {
		return nil, err
	}
	u := r.Store.User().ModelFromDTO(userDTO)

	return &model.Employee{
		EmployeeID: dto.EmployeeID,
		User:       *u,
		Hotel:      *hotel,
		Position:   model.Position(dto.Position),
	}, nil
}
