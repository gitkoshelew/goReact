package store

import (
	"errors"
	"hotel/domain/model"
	"log"
)

// EmployeeRepository ...
type EmployeeRepository struct {
	Store *Store
}

// Create employee and save it to DB
func (r *EmployeeRepository) Create(e *model.Employee) (*model.Employee, error) {
	if err := r.Store.Db.QueryRow(`INSERT INTO employee 
	(email, verified, role, first_name, surname, middle_name, sex, 
	date_of_birth, address, phone, photo, hotel_id, position) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`,
		e.Email,
		e.Verified,
		e.Role,
		e.Name,
		e.Surname,
		e.MiddleName,
		e.Sex,
		e.DateOfBirth,
		e.Address,
		e.Phone,
		e.Photo,
		e.Hotel.HotelID,
		e.Position,
	).Scan(&e.EmployeeID); err != nil {
		r.Store.Logger.Errorf("Can't create employee. Err msg:%v.", err)
		return nil, err
	}
	r.Store.Logger.Info("Creat employee with id = %d", e.EmployeeID)
	return e, nil
}

// GetAll returns all employees
func (r *EmployeeRepository) GetAll() (*[]model.EmployeeDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM employee")
	if err != nil {
		r.Store.Logger.Errorf("Can't find employees. Err msg: %v", err)
	}
	employees := []model.EmployeeDTO{}

	for rows.Next() {
		employee := model.EmployeeDTO{}
		err := rows.Scan(
			&employee.EmployeeID,
			&employee.Email,
			&employee.Verified,
			&employee.Role,
			&employee.Name,
			&employee.Surname,
			&employee.MiddleName,
			&employee.Sex,
			&employee.DateOfBirth,
			&employee.Address,
			&employee.Phone,
			&employee.Photo,
			&employee.HotelID,
			&employee.Position,
		)
		if err != nil {
			r.Store.Logger.Errorf("Can't find employees. Err msg: %v", err)
			continue
		}
		employees = append(employees, employee)
	}
	return &employees, nil
}

// FindByID searchs and returns employee by ID
func (r *EmployeeRepository) FindByID(id int) (*model.EmployeeDTO, error) {
	employee := &model.EmployeeDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM employee WHERE id = $1", id).Scan(
		&employee.EmployeeID,
		&employee.Email,
		&employee.Verified,
		&employee.Role,
		&employee.Name,
		&employee.Surname,
		&employee.MiddleName,
		&employee.Sex,
		&employee.DateOfBirth,
		&employee.Address,
		&employee.Phone,
		&employee.Photo,
		&employee.HotelID,
		&employee.Position,
	); err != nil {
		r.Store.Logger.Errorf("Cant find employee. Err msg:%v.", err)
		return nil, err
	}
	return employee, nil
}

// Delete employee from DB by ID
func (r *EmployeeRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM employee WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Can't delete employee. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Can't delete employee. Err msg:%v.", err)

		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Can't delete employee. Err msg:%v.", err)
		return err
	}

	log.Printf("Employee deleted, rows affectet: %d", result)
	return nil
}

// Update employee from DB
func (r *EmployeeRepository) Update(e *model.Employee) error {

	result, err := r.Store.Db.Exec(
		`UPDATE employee SET 
		email = $1, 
		verified = $2, 
		role = $3, 
		first_name = $4, 
		surname = $5, 
		middle_name = $6, 
		sex = $7, 
		date_of_birth = $8, 
		address = $9, 
		phone = $10, 
		photo = $11, 
		hotel_id = $12,
		position = $13
		WHERE id = $14`,
		e.Email,
		e.Verified,
		e.Role,
		e.Name,
		e.Surname,
		e.MiddleName,
		e.Sex,
		e.DateOfBirth,
		e.Address,
		e.Phone,
		e.Photo,
		e.Hotel.HotelID,
		e.Position,
		e.EmployeeID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Can't update employee. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Update employee with id = %d,rows affectet: %d ", e.EmployeeID, result)
	return nil
}

// EmployeeFromDTO ...
func (r *EmployeeRepository) EmployeeFromDTO(dto *model.EmployeeDTO) (*model.Employee, error) {
	hotel, err := r.Store.Hotel().FindByID(dto.HotelID)
	if err != nil {
		r.Store.Logger.Errorf("Can't convert employeeDTO. Err msg: %v", err)
		return nil, err
	}
	return &model.Employee{
		EmployeeID:  dto.EmployeeID,
		Email:       dto.Email,
		Role:        model.Role(dto.Role),
		Verified:    dto.Verified,
		Name:        dto.Name,
		Surname:     dto.Surname,
		MiddleName:  dto.MiddleName,
		Sex:         dto.Sex,
		DateOfBirth: dto.DateOfBirth,
		Address:     dto.Address,
		Phone:       dto.Phone,
		Photo:       dto.Photo,
		Hotel:       *hotel,
		Position:    model.Position(dto.Position),
	}, nil

}
