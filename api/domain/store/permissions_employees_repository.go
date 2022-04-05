package store

import (
	"errors"
	"goReact/domain/model"
)

// PermissionsEmployeeRepository serves to communicate employees and permissions
type PermissionsEmployeeRepository struct {
	Store *Store
}

// GetAll ...
func (r *PermissionsEmployeeRepository) GetAll() (*[]model.PermissionsEmployeesDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM permissions_employees")
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while getting all permissions for employees. Err msg: %v", err)
	}

	permissionsEmployees := []model.PermissionsEmployeesDTO{}

	for rows.Next() {
		permissionEmployees := model.PermissionsEmployeesDTO{}
		err := rows.Scan(
			&permissionEmployees.PermissionsID,
			&permissionEmployees.EmployeeID,
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
func (r *PermissionsEmployeeRepository) SetForEmployee(pe *model.PermissionsEmployees) error {

	result, err := r.Store.Db.Exec("INSERT INTO permissions_employees (permissions_id, employee_id) VALUES ($1, $2)", pe.Permissions.PermissionID, pe.Employee.EmployeeID)
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

// ModelFromDTO ...
func (r *PermissionsEmployeeRepository) ModelFromDTO(dto *model.PermissionsEmployeesDTO) (*model.PermissionsEmployees, error) {
	per, err := r.Store.Permissions().FindByID(dto.PermissionsID)
	if err != nil {
		return nil, err
	}
	employeeDTO, err := r.Store.Employee().FindByID(dto.EmployeeID)
	if err != nil {
		return nil, err
	}
	employee, err := r.Store.Employee().ModelFromDTO(employeeDTO)
	if err != nil {
		return nil, err
	}

	return &model.PermissionsEmployees{
		Permissions: *per,
		Employee:    *employee,
	}, nil
}
