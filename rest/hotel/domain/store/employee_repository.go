package store

import (
	"fmt"
	"hotel/domain/model"
	"strings"
	"time"
)

// EmployeeRepository ...
type EmployeeRepository struct {
	Store *Store
}

// Create employee and save it to DB
func (r *EmployeeRepository) Create(e *model.Employee) (*int, error) {
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

	r.Store.Logger.Infof("Creat employee with id = %d", e.EmployeeID)

	return &e.EmployeeID, nil
}

// GetAll returns all employees
func (r *EmployeeRepository) GetAll() (*[]model.EmployeeDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM employee")
	if err != nil {
		r.Store.Logger.Errorf("Can't find employees. Err msg: %v", err)
		return nil, err
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
			r.Store.Logger.Debugf("Can't find employees. Err msg: %v", err)
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
		r.Store.Logger.Errorf("Can't delete employee. Err msg:%v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Employee with id %d was deleted", id)
	return nil
}

// FindByEmail searchs and returns user by email
func (r *EmployeeRepository) FindByEmail(email string) (*model.EmployeeDTO, error) {
	employee := &model.EmployeeDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM employee WHERE email = $1",
		email).Scan(
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
		r.Store.Logger.Errorf("Eror occure while searching user by email. Err msg: %v", err)
		return nil, err
	}
	return employee, nil
}

// Update employee from DB
func (r *EmployeeRepository) Update(e *model.Employee) error {
	email := "email"
	if e.Email != "" {
		var emailIsUsed bool
		err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT email FROM employee WHERE email = $1)", e.Email).Scan(&emailIsUsed)
		if err != nil {
			r.Store.Logger.Errorf("Eror during checking employees email or password. Err msg: %v", err)
			return err
		}

		if emailIsUsed {
			r.Store.Logger.Errorf("email is used. Err msg: %v", ErrEmailIsUsed)
			return ErrEmailIsUsed
		}

		email = fmt.Sprintf("'%s'", e.Email)
	}

	role := "role"
	if e.Role != "" {
		role = fmt.Sprintf("'%s'", string(e.Role))
	}
	verified := "verified"
	if e.Verified != nil {
		verified = fmt.Sprintf("%v", *e.Verified)
	}
	name := "first_name"
	if e.Name != "" {
		name = fmt.Sprintf("'%s'", e.Name)
	}
	surname := "surname"
	if e.Surname != "" {
		surname = fmt.Sprintf("'%s'", e.Surname)
	}
	middlename := "middle_name"
	if e.MiddleName != "" {
		middlename = fmt.Sprintf("'%s'", e.MiddleName)
	}
	sex := "sex"
	if e.Sex != "" {
		sex = fmt.Sprintf("'%s'", string(e.Sex))
	}
	dateOfBirth := "date_of_birth"
	if e.DateOfBirth != nil {
		dateOfBirth = fmt.Sprintf("'%s'", e.DateOfBirth.Format(time.RFC3339))
	}
	address := "address"
	if e.Address != "" {
		address = fmt.Sprintf("'%s'", e.Address)
	}
	phone := "phone"
	if e.Phone != "" {
		phone = fmt.Sprintf("'%s'", e.Phone)
	}
	photo := "photo"
	if e.Photo != "" {
		photo = fmt.Sprintf("'%s'", e.Photo)
	}
	hotelID := " hotel_id"
	if e.Hotel.HotelID != 0 {
		hotelID = fmt.Sprintf("%d", e.Hotel.HotelID)
	}
	position := "position"
	if e.Position != "" {
		position = fmt.Sprintf("'%s'", string(e.Position))
	}

	result, err := r.Store.Db.Exec(fmt.Sprintf(
		`UPDATE employee SET 
			email = %s, role = %s, verified = %s, 
			first_name = %s, surname = %s, middle_name = %s, sex = %s, date_of_birth = %s, 
			address = %s, phone = %s, photo = %s, hotel_id = %s, position = %s 
			WHERE id = $1`,
		email,
		role,
		verified,
		strings.Title(strings.ToLower(name)),
		strings.Title(strings.ToLower(surname)),
		strings.Title(strings.ToLower(middlename)),
		sex,
		dateOfBirth,
		address,
		phone,
		photo,
		hotelID,
		position,
	), e.EmployeeID)
	if err != nil {
		r.Store.Logger.Errorf("Erorr occured while updating user. Err msg: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating user. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while updating user. Err msg:%v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("User with id %d was updated", e.EmployeeID)

	return nil
}

// EmailCheck check if email exists in DB
func (r *EmployeeRepository) EmailCheck(email string) (*bool, error) {
	var emailIsUsed bool
	err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT email FROM employee WHERE email = $1)", email).Scan(&emailIsUsed)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while email checking. Err msg: %v", err)
		return &emailIsUsed, err
	}
	return &emailIsUsed, nil
}

// ModelFromDTO ...
func (r *EmployeeRepository) ModelFromDTO(dto *model.EmployeeDTO) (*model.Employee, error) {
	hotel, err := r.Store.Hotel().FindByID(dto.HotelID)
	if err != nil {
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
		Sex:         model.Sex(dto.Sex),
		DateOfBirth: dto.DateOfBirth,
		Address:     dto.Address,
		Phone:       dto.Phone,
		Photo:       dto.Photo,
		Hotel:       *hotel,
		Position:    model.Position(dto.Position),
	}, nil

}
