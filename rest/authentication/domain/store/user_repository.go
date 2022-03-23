package store

import (
	"auth/domain/model"
	"auth/domain/utils"
	"fmt"
)

// UserRepository ...
type UserRepository struct {
	Store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) (*int, error) {

	var emailIsUsed bool
	err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT email FROM users WHERE email = $1)", u.Email).Scan(&emailIsUsed)
	if err != nil {
		r.Store.Logger.Errorf("Eror during checking users email or password. Err msg: %v", err)
		return nil, err
	}

	if emailIsUsed {
		r.Store.Logger.Errorf("email is used. Err msg: %v", ErrEmailIsUsed)
		return nil, ErrEmailIsUsed
	}

	if err := r.Store.Db.QueryRow(
		`INSERT INTO users 
		(email, password, role, verified) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`,
		u.Email,
		u.Password,
		string(u.Role),
		false,
	).Scan(&u.UserID); err != nil {
		r.Store.Logger.Errorf("Eror occured while creating user. Err msg: %v", err)
		return nil, err
	}
	return &u.UserID, nil
}

// GetAll returns all users
func (r *UserRepository) GetAll() (*[]model.User, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM users")
	if err != nil {
		r.Store.Logger.Errorf("Eror occured while getting all users. Err msg: %v", err)
		return nil, err
	}
	users := []model.User{}

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(
			&user.UserID,
			&user.Email,
			&user.Password,
			&user.Verified,
			&user.Role,
		)
		if err != nil {
			r.Store.Logger.Debugf("Eror occured while getting all bookings. Err msg: %v", err)
			continue
		}
		users = append(users, user)
	}
	return &users, nil
}

// FindByEmail searchs and returns user by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := r.Store.Db.QueryRow("SELECT * FROM users WHERE email = $1",
		email).Scan(
		&user.UserID,
		&user.Email,
		&user.Password,
		&user.Verified,
		&user.Role,
	); err != nil {
		r.Store.Logger.Errorf("Eror during checking users email or password. Err msg: %v", err)
		return nil, err
	}
	return user, nil
}

// FindByID searchs and returns user by ID
func (r *UserRepository) FindByID(id int) (*model.UserDTO, error) {
	user := &model.UserDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM users WHERE id = $1",
		id).Scan(
		&user.UserID,
		&user.Email,
		&user.Password,
		&user.Verified,
		&user.Role,
	); err != nil {
		r.Store.Logger.Errorf("Eror occure while searching user by id. Err msg: %v", err)
		return nil, err
	}
	return user, nil
}

// Delete user from DB by ID
func (r *UserRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting user. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting user. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while deleting user. Err msg:%v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("User with id %d was deleted", id)
	return nil
}

// Update user from DB
func (r *UserRepository) Update(u *model.User) error {

	email := "email"
	if u.Email != "" {
		var emailIsUsed bool
		err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT email FROM users WHERE email = $1)", u.Email).Scan(&emailIsUsed)
		if err != nil {
			r.Store.Logger.Errorf("Eror during checking users email or password. Err msg: %v", err)
			return err
		}

		if emailIsUsed {
			r.Store.Logger.Errorf("email is used. Err msg: %v", ErrEmailIsUsed)
			return ErrEmailIsUsed
		}

		email = fmt.Sprintf("'%s'", u.Email)
	}

	password := "password"
	if u.Password != "" {
		encryptedPassword, err := utils.EncryptPassword(u.Password)
		if err != nil {
			r.Store.Logger.Errorf("Cant't encrypt password. Err msg: %v", err)
			return err
		}
		password = fmt.Sprintf("'%s'", encryptedPassword)
	}

	role := "role"
	if u.Role != "" {
		role = fmt.Sprintf("'%s'", string(u.Role))
	}
	verified := "verified"
	if u.Verified != nil {
		verified = fmt.Sprintf("%v", *u.Verified)
	}

	result, err := r.Store.Db.Exec(fmt.Sprintf(
		`UPDATE users SET 
			email = %s, password = %s, role = %s, verified = %s 
			WHERE id = $1`,
		email,
		password,
		role,
		verified,
	), u.UserID)
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

	r.Store.Logger.Infof("User with id %d was updated", u.UserID)

	return nil
}

// VerifyEmail user from DB
func (r *UserRepository) VerifyEmail(userID int) error {
	result, err := r.Store.Db.Exec(
		"UPDATE users SET verified = $1, role = $2  WHERE id = $3",
		true,
		string(model.ClientRole),
		userID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while verifying user. Err msg:%v.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while verifying user. Err msg:%v.", err)
		return nil
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("No rows was affected, possible reason: no user with such ID, err msg: %v", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}
	return nil
}

// ModelFromDTO ...
func (r *UserRepository) ModelFromDTO(u *model.UserDTO) (*model.User, error) {
	return &model.User{
		UserID:   u.UserID,
		Email:    u.Email,
		Password: u.Password,
		Role:     model.Role(u.Role),
		Verified: u.Verified,
	}, nil
}
