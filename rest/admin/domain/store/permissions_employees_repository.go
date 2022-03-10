package store

import (
	"admin/domain/model"
	"errors"
)

// PermissionsEmployeeRepository serves to communicate employees and permissions
type PermissionsEmployeeRepository struct {
	Store *Store
}

// GetAll ...
func (r *PermissionsEmployeeRepository) GetAll() (*[]model.PermissionsEmployees, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM permissions_employees")
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while getting all permissions for employees. Err msg: %v", err)
	}

	permissionsEmployees := []model.PermissionsEmployees{}

	for rows.Next() {
		permissionEmployees := model.PermissionsEmployees{}
		err := rows.Scan(
			&permissionEmployees.Permissions.PermissionID,
			&permissionEmployees.Employee.EmployeeID,
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occurred while getting all permissions for employees. Err msg: %v", err)
			continue
		}
		permissionsEmployees = append(permissionsEmployees, permissionEmployees)
	}
	return &permissionsEmployees, nil
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
		r.Store.Logger.Errorf("Error occurred while setting permissions for employees. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Error occurred while setting permissions for employees. Err msg:%v.", err)
		return err
	}

	r.Store.Logger.Info("Permissions for employees seted: %d", result)
	return nil

}
