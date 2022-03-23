package store

import (
	"auth/domain/model"
	"errors"
)

// UserRepository ...
type UserRepository struct {
	Store *Store
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
		r.Store.Logger.Errorf("Eror during checking users email or password. Err msg: %w", err)
		return nil, err
	}
	return user, nil
}

// Create ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {

	var emailIsUsed bool
	err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT email FROM users WHERE email = $1)", u.Email).Scan(&emailIsUsed)
	if err != nil {
		r.Store.Logger.Errorf("Eror during checking users email or password. Err msg: %w", err)
		return nil, err
	}

	if emailIsUsed {
		r.Store.Logger.Errorf("email is used. Err msg: %w", err)
		return nil, errors.New("Email already in use")
	}

	if err := r.Store.Db.QueryRow(
		`INSERT INTO users 
		(email, password, role, verified) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`,
		u.Email,
		u.Password,
		string(u.Role),
		u.Verified,
	).Scan(&u.UserID); err != nil {
		r.Store.Logger.Errorf("Eror occured while creating user. Err msg: %w", err)
		return nil, err
	}
	return u, nil
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
