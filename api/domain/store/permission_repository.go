package store

import (
	"goReact/domain/model"
	"errors"
)

// PermissionsRepository ...
type PermissionsRepository struct {
	Store *Store
}

// Create permossion and save it to DB
func (r *PermissionsRepository) Create(p *model.Permission) (*int, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO PERMISSIONS (name, description) VALUES ($1, $2) RETURNING permissions_id",
		p.Name,
		p.Descriptoin,
	).Scan(&p.PermissionID); err != nil {
		r.Store.Logger.Errorf("Error occurred while creating permission. Err msg:%v.", err)
		return nil, err
	}
	return &p.PermissionID, nil
}

// GetAll ...
func (r *PermissionsRepository) GetAll() (*[]model.Permission, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM PERMISSIONS")
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while getting all permissions. Err msg: %v", err)
	}

	permissoins := []model.Permission{}

	for rows.Next() {
		permissoin := model.Permission{}
		err := rows.Scan(
			&permissoin.PermissionID,
			&permissoin.Name,
			&permissoin.Descriptoin,
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occurred while getting all permissions. Err msg: %v", err)
			continue
		}
		permissoins = append(permissoins, permissoin)
	}
	return &permissoins, nil
}

// FindByID ...
func (r *PermissionsRepository) FindByID(id int) (*model.Permission, error) {

	permissoin := model.Permission{}
	if err := r.Store.Db.QueryRow("SELECT * FROM PERMISSIONS WHERE permissions_id = $1", id).Scan(
		&permissoin.PermissionID,
		&permissoin.Name,
		&permissoin.Descriptoin,
	); err != nil {
		r.Store.Logger.Errorf("Error occurred while getting permission by id. Err msg: %v", err)
		return nil, err
	}
	return &permissoin, nil
}

// Delete ...
func (r *PermissionsRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM PERMISSIONS WHERE permissions_id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while deleting permission. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while deleting permission. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Error occurred while deleting permission. Err msg:%v.", err)
		return err
	}

	r.Store.Logger.Info("Permission deleted, rows affectet: %d", result)
	return nil
}

// GetEmployeeByID return []permissions that the employee has
func (r *PermissionsRepository) GetEmployeeByID(id int) (*[]model.Permission, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM PERMISSiONS WHERE permissions_id IN ( SELECT permissions_id FROM permissions_employees where employee_id = $1 )", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while getting permissions. Err msg: %v", err)
	}

	permissoins := []model.Permission{}

	for rows.Next() {
		permissoin := model.Permission{}
		err := rows.Scan(
			&permissoin.PermissionID,
			&permissoin.Name,
			&permissoin.Descriptoin,
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occurred while getting permissions by employee id. Err msg: %v", err)
			continue
		}
		permissoins = append(permissoins, permissoin)
	}
	return &permissoins, nil
}
