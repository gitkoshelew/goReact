package store

import (
	"admin/domain/model"
	"errors"
)

// PermissionsEmployeeRepository serves to communicate employees and permissions
type PermissionsEmployeeRepository struct {
	Store *Store
}

func (r *PermissionsEmployeeRepository) GetAll() (*[]model.Permissions_employees, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM permissions_employees")
	if err != nil {
		r.Store.Logger.Errorf("Can't find Permissions for employees. Err msg: %v", err)
	}

	permissions_employees := []model.Permissions_employees{}

	for rows.Next() {
		permission_employees := model.Permissions_employees{}
		err := rows.Scan(
			&permission_employees.Permissions.PermissionID,
			&permission_employees.Employee.EmployeeID,
		)
		if err != nil {
			r.Store.Logger.Errorf("Can't find permissions_employees. Err msg: %v", err)
			continue
		}
		permissions_employees = append(permissions_employees, permission_employees)
	}
	return &permissions_employees, nil
}

//SetForEmployee set permissions for employee
func (r *PermissionsEmployeeRepository) SetForEmployee(PermissionID int, employeeID int) error {

	result, err := r.Store.Db.Exec("INSERT INTO permissions_employees (permissions_id, employee_id) VALUES ($1, $2)", PermissionID, employeeID)
	if err != nil {
		r.Store.Logger.Info("Permissions for employees seted: %d", result)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Can't set Permissions for employees. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Can't set Permissions for employees. Err msg:%v.", err)
		return err
	}

	r.Store.Logger.Info("Permissions for employees seted: %d", result)
	return nil

}
