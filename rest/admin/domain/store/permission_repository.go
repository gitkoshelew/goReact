package store

import (
	"admin/domain/model"
	"errors"
)

// PermissionsRepository ...
type PermissionsRepository struct {
	Store *Store
}

// Create permossion and save it to DB
func (r *PermissionsRepository) Create(p *model.Permission) (*model.Permission, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO PERMISSIONS (name, description) VALUES ($1, $2) RETURNING id",
		p.Name,
		p.Descriptoin,
	).Scan(&p.PermissionID); err != nil {
		r.Store.Logger.Errorf("Can't create permission. Err msg:%v.", err)
		return nil, err
	}
	return p, nil
}

func (r *PermissionsRepository) GetAll() (*[]model.Permission, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM PERMISSIONS")
	if err != nil {
		r.Store.Logger.Errorf("Can't find permissions. Err msg: %v", err)
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
			r.Store.Logger.Errorf("Can't find permissions. Err msg: %v", err)
			continue
		}
		permissoins = append(permissoins, permissoin)
	}
	return &permissoins, nil
}

func (r *PermissionsRepository) FindByID(id int) (*model.Permission, error) {

	permissoin := model.Permission{}
	if err := r.Store.Db.QueryRow("SELECT * FROM PERMISSIONS WHERE id = $1", id).Scan(
		&permissoin.PermissionID,
		&permissoin.Name,
		&permissoin.Descriptoin,
	); err != nil {
		r.Store.Logger.Errorf("Can't find permissions. Err msg: %v", err)
		return nil, err
	}
	return &permissoin, nil
}

func (r *PermissionsRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM PERMISSIONS WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Can't delete permission. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Can't delete permission. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Can't delete permission. Err msg:%v.", err)
		return err
	}

	r.Store.Logger.Info("Permission deleted, rows affectet: %d", result)
	return nil
}

// GetByEmployeeId return []permissions that the employee has
func (r *PermissionsRepository) GetByEmployeeId(id int) (*[]model.Permission, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM PERMISSiONS WHERE id IN ( SELECT permissions_id FROM permissions_employees where employee_id = $1 )", id)
	if err != nil {
		r.Store.Logger.Errorf("Can't find permissions. Err msg: %v", err)
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
			r.Store.Logger.Errorf("Can't find permissions. Err msg: %v", err)
			continue
		}
		permissoins = append(permissoins, permissoin)
	}
	return &permissoins, nil
}
