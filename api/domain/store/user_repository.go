package store

import (
	"errors"
	"goReact/domain/model"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository ...
type UserRepository struct {
	Store *Store
}

var (
	// ErrEmailIsUsed ...
	ErrEmailIsUsed = errors.New("Email already in use")
)

// Create user and save it to DB
func (r *UserRepository) Create(u *model.UserDTO) (*model.UserDTO, error) {

	if err := u.ModelFromDTO().Validate(); err != nil {
		r.Store.Logger.Errorf("Error occured while validating user. Err msg: %w", err)
		return nil, err
	}

	var emailIsUsed bool
	err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT email FROM users WHERE email = $1)", u.Email).Scan(&emailIsUsed)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while email validating. Err msg: %w", err)
		return nil, err
	}

	if emailIsUsed {
		r.Store.Logger.Error(ErrEmailIsUsed)
		return nil, ErrEmailIsUsed
	}

	r.Store.EncryptPassword(&u.Password)

	if err := r.Store.Db.QueryRow(
		`INSERT INTO users 
		(email, password, role, verified, first_name, surname, middle_name, sex, date_of_birth, address, phone, photo) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
		RETURNING id`,
		u.Email,
		u.Password,
		u.Role,
		u.Verified,
		u.Name,
		u.Surname,
		u.MiddleName,
		u.Sex,
		u.DateOfBirth,
		u.Address,
		u.Phone,
		u.Photo,
	).Scan(&u.UserID); err != nil {
		r.Store.Logger.Errorf("Error occured while inserting data to DB. Err msg: %w", err)
		return nil, err
	}
	return u, nil
}

// GetAll returns all users
func (r *UserRepository) GetAll() (*[]model.User, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	users := []model.User{}

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(
			&user.UserID,
			&user.Name,
			&user.Surname,
			&user.MiddleName,
			&user.Email,
			&user.DateOfBirth,
			&user.Address,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Verified,
			&user.Sex,
			&user.Photo,
		)
		if err != nil {
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
		&user.Name,
		&user.Surname,
		&user.MiddleName,
		&user.Email,
		&user.DateOfBirth,
		&user.Address,
		&user.Phone,
		&user.Password,
		&user.Role,
		&user.Verified,
		&user.Sex,
		&user.Photo,
	); err != nil {
		r.Store.Logger.Errorf("Eror during checking users email or password. Err msg: %s", err.Error())
		return nil, err
	}
	return user, nil
}

// FindByID searchs and returns user by ID
func (r *UserRepository) FindByID(id int) (*model.User, error) {
	user := &model.User{}
	if err := r.Store.Db.QueryRow("SELECT * FROM users WHERE id = $1",
		id).Scan(
		&user.UserID,
		&user.Name,
		&user.Surname,
		&user.MiddleName,
		&user.Email,
		&user.DateOfBirth,
		&user.Address,
		&user.Phone,
		&user.Password,
		&user.Role,
		&user.Verified,
		&user.Sex,
		&user.Photo,
	); err != nil {
		r.Store.Logger.Errorf("Cant find user. Err msg:%v.", err)
		return nil, err
	}
	return user, nil
}

// Delete user from DB by ID
func (r *UserRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("No rows affected")
	}
	return nil
}

// Update user from DB
func (r *UserRepository) Update(u *model.User) error {
	encryptedPassword, err := model.EncryptPassword(u.Password)
	if err != nil {
		r.Store.Logger.Errorf("Cant't encrypt password. Err msg: %v", err)
		return err
	}

	_, err = r.Store.Db.Exec(
		`UPDATE users SET 
			email = $1, password = $2, role = $3, verified = $4, 
			first_name = $5, surname = $6, middle_name = $7, sex = $8, date_of_birth = $9, 
			address = $10, phone = $11, photo = $12 
			WHERE id = $13`,
		u.Email,
		encryptedPassword,
		string(u.Role),
		u.Verified,
		u.Name,
		u.Surname,
		u.MiddleName,
		string(u.Sex),
		u.DateOfBirth,
		u.Address,
		u.Phone,
		u.Photo,
		u.UserID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Cant't set into users table. Err msg: %v", err)
		return err
	}

	return nil
}

// VerifyEmail user from DB
func (r *UserRepository) VerifyEmail(userID int) error {
	result, err := r.Store.Db.Exec(
		"UPDATE users SET verified = $1 WHERE id = $2",
		true,
		userID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating user. Err msg:%v.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating user. Err msg:%v.", err)
		return nil
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("No rows was affected, possible reason: no user with such ID. Err msg:%v.", err)
	}
	return nil
}

// EmailCheck check if email exists in DB
func (r *UserRepository) EmailCheck(email string) *bool {
	var emailIsUsed bool
	err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT email FROM users WHERE email = $1)", email).Scan(&emailIsUsed)
	if err != nil {
		r.Store.Logger.Errorf("Error while checking if email ia already in use. Err msg: %v", err)
	}
	return &emailIsUsed
}

// PasswordChange ...
func (r *UserRepository) PasswordChange(u *model.User) error {
	encryptedPassword, err := model.EncryptPassword(u.Password)
	if err != nil {
		r.Store.Logger.Errorf("Cant't encrypt password. Err msg: %v", err)
		return err
	}

	result, err := r.Store.Db.Exec("UPDATE users SET password = $1 WHERE id = $2", encryptedPassword, u.UserID)
	if err != nil {
		r.Store.Logger.Errorf("Cant't set into users table. Err msg: %v", err)
		return err
	}
	r.Store.Logger.Infof("User updated, rows affectet: %d", result)
	return nil
}

// CheckPasswordHash if passwords are same err=nil
func (r *UserRepository) CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	r.Store.Logger.Infof("Eror during checking users email or password. Err msg: %s", err.Error())
	return err
}
