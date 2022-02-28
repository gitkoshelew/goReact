package store

import (
	"admin/domain/model"
	"errors"
	"log"
)

// PermissionsEmployeeRepository serves to communicate employees and permissions
type PermissionsEmployeeRepository struct {
	Store *Store
}

func (r *PermissionsEmployeeRepository) GetAll() (*[]model.Permissions_employees, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM permissions_employees")
	if err != nil {
		log.Print(err)
	}

	permissions_employees := []model.Permissions_employees{}

	for rows.Next() {
		permission_employees := model.Permissions_employees{}
		err := rows.Scan(
			&permission_employees.Permissions.PermissionID,
			&permission_employees.Employee.EmployeeID,
		)
		if err != nil {
			log.Print(err)
			continue
		}
		permissions_employees = append(permissions_employees, permission_employees)
	}
	return &permissions_employees, nil
}

func (r *PermissionsEmployeeRepository) SetForEmployee(PermissionID int, employeeID int) error {

	result, err := r.Store.Db.Exec("INSERT INTO permissions_employees (permissions_id, employee_id) VALUES ($1, $2)", PermissionID, employeeID)
	if err != nil {
		log.Print(err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("No rows affected")
	}

	log.Printf("Permissoin seted ")
	return nil

}
