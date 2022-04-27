package store

import (
	"fmt"
	"goReact/domain/model"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository ...
type UserRepository struct {
	Store *Store
}

// Create user from oauth and save it to DB
func (r *UserRepository) CreateFromSocial(user *model.User) (*int, error) {

	if user.Email != "" {
		emailIsUsed, err := r.EmailCheck(user.Email)
		if err != nil {
			r.Store.Logger.Errorf("Error occured while email validating. Err msg: %v", err)
			return nil, err
		}
		if *emailIsUsed {
			r.Store.Logger.Error(ErrEmailIsUsed)
			return nil, ErrEmailIsUsed
		}
	}
	
	socialNetworkId, err := r.CheckSocialNetworkID(user.SocialNetworkID)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while social network id validating. Err msg: %v", err)
		return nil, err
	}

	if *socialNetworkId {
		r.Store.Logger.Error(ErrSocialIDIsExist)
		return nil, ErrSocialIDIsExist
	}

	if err := r.Store.Db.QueryRow(
		`INSERT INTO users 
		(email,  role, verified, first_name, surname, middle_name, sex, date_of_birth, address, phone, photo , social_network, social_network_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 , $12 , $13) 
		RETURNING user_id`,
		user.Email,
		string(model.ClientRole),
		false,
		strings.Title(strings.ToLower(user.Name)),
		strings.Title(strings.ToLower(user.Surname)),
		strings.Title(strings.ToLower(user.MiddleName)),
		string(user.Sex),
		user.DateOfBirth,
		user.Address,
		user.Phone,
		user.Photo,
		user.SocialNetwork,
		user.SocialNetworkID,
	).Scan(&user.UserID); err != nil {
		r.Store.Logger.Errorf("Error occured while inserting data to DB. Err msg: %v", err)
		return nil, err
	}

	r.Store.Logger.Infof("User with id %d was created.", user.UserID)

	return &user.UserID, nil
}

// Create user and save it to DB
func (r *UserRepository) Create(user *model.User) (*int, error) {

	var emailIsUsed bool
	err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT email FROM users WHERE email = $1)", user.Email).Scan(&emailIsUsed)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while email validating. Err msg: %v", err)
		return nil, err
	}

	if emailIsUsed {
		r.Store.Logger.Error(ErrEmailIsUsed)
		return nil, ErrEmailIsUsed
	}
	encryptedPassword, err := model.EncryptPassword(user.Password)
	if err != nil {
		r.Store.Logger.Error("Eror occured while password encrypting. Err msg: %v", err)
		return nil, err
	}

	user.Password = encryptedPassword

	if err := r.Store.Db.QueryRow(
		`INSERT INTO users 
		(email, password, role, verified, first_name, surname, middle_name, sex, date_of_birth, address, phone, photo) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
		RETURNING user_id`,
		user.Email,
		user.Password,
		string(model.AnonymousRole),
		false,
		strings.Title(strings.ToLower(user.Name)),
		strings.Title(strings.ToLower(user.Surname)),
		strings.Title(strings.ToLower(user.MiddleName)),
		string(user.Sex),
		user.DateOfBirth,
		user.Address,
		user.Phone,
		user.Photo,
	).Scan(&user.UserID); err != nil {
		r.Store.Logger.Errorf("Error occured while inserting data to DB. Err msg: %v", err)
		return nil, err
	}

	r.Store.Logger.Infof("User with id %d was created.", user.UserID)

	return &user.UserID, nil
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
			&user.SocialNetwork,
			&user.SocialNetworkID,
		)
		if err != nil {
			r.Store.Logger.Debugf("Eror occured while getting all user. Err msg: %v", err)
			continue
		}
		users = append(users, user)
	}
	return &users, nil
}

// FindByEmail searchs and returns user by email
func (r *UserRepository) FindByEmail(email string) (*model.UserDTO, error) {
	user := &model.UserDTO{}
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
		&user.SocialNetwork,
		&user.SocialNetworkID,
	); err != nil {
		r.Store.Logger.Errorf("Eror occure while searching user by email. Err msg: %v", err)
		return nil, err
	}
	return user, nil
}

// FindByID searchs and returns user by ID
func (r *UserRepository) FindByID(id int) (*model.UserDTO, error) {
	user := &model.UserDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM users WHERE user_id = $1",
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
		&user.SocialNetwork,
		&user.SocialNetworkID,
	); err != nil {
		r.Store.Logger.Errorf("Eror occure while searching user by id. Err msg: %v", err)
		return nil, err
	}
	return user, nil
}

// Delete user from DB by ID
func (r *UserRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM users WHERE user_id = $1", id)
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
		encryptedPassword, err := model.EncryptPassword(u.Password)
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
	name := "first_name"
	if u.Name != "" {
		name = fmt.Sprintf("'%s'", u.Name)
	}
	surname := "surname"
	if u.Surname != "" {
		surname = fmt.Sprintf("'%s'", u.Surname)
	}
	middlename := "middle_name"
	if u.MiddleName != "" {
		middlename = fmt.Sprintf("'%s'", u.MiddleName)
	}
	sex := "sex"
	if u.Sex != "" {
		sex = fmt.Sprintf("'%s'", string(u.Sex))
	}
	dateOfBirth := "date_of_birth"
	if u.DateOfBirth != nil {
		dateOfBirth = fmt.Sprintf("'%s'", u.DateOfBirth.Format(time.RFC3339))
	}
	address := "address"
	if u.Address != "" {
		address = fmt.Sprintf("'%s'", u.Address)
	}
	phone := "phone"
	if u.Phone != "" {
		phone = fmt.Sprintf("'%s'", u.Phone)
	}
	photo := "photo"
	if u.Photo != "" {
		photo = fmt.Sprintf("'%s'", u.Photo)
	}
	socialNetwork := "social_network"
	if u.SocialNetwork != "" {
		socialNetwork = fmt.Sprintf("'%s'", string(u.SocialNetwork))
	}
	socialNetworkID := "social_network_id"
	if u.SocialNetworkID != "" {
		socialNetworkID = fmt.Sprintf("'%s'", u.SocialNetworkID)
	}

	result, err := r.Store.Db.Exec(fmt.Sprintf(
		`UPDATE users SET 
			email = %s, password = %s, role = %s, verified = %s, 
			first_name = %s, surname = %s, middle_name = %s, sex = %s, date_of_birth = %s, 
			address = %s, phone = %s, photo = %s  , social_network = %s ,social_network_id = %s
			WHERE user_id = $1`,
		email,
		password,
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
		socialNetwork,
		socialNetworkID,
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
		"UPDATE users SET verified = $1, role = $2  WHERE user_id = $3",
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

// EmailCheck check if email exists in DB
func (r *UserRepository) EmailCheck(email string) (*bool, error) {
	var emailIsUsed bool
	err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT email FROM users WHERE email = $1)", email).Scan(&emailIsUsed)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while email checking. Err msg: %v", err)
		return &emailIsUsed, err
	}
	return &emailIsUsed, nil
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

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while verifying user. Err msg:%v.", err)
		return nil
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("No rows was affected, possible reason: no user with such ID, err msg: %v", ErrNoRowsAffected)
		return ErrNoRowsAffected
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

// ModelFromDTO ...
func (r *UserRepository) ModelFromDTO(u *model.UserDTO) (*model.User, error) {
	return &model.User{
		UserID:          u.UserID,
		Email:           u.Email,
		Password:        u.Password,
		Role:            model.Role(u.Role),
		Verified:        u.Verified,
		Name:            u.Name,
		Surname:         u.Surname,
		MiddleName:      u.MiddleName,
		Sex:             model.Sex(u.Sex),
		DateOfBirth:     u.DateOfBirth,
		Address:         u.Address,
		Phone:           u.Phone,
		Photo:           u.Photo,
		SocialNetwork:   model.SocialNetwork(u.SocialNetwork),
		SocialNetworkID: u.SocialNetworkID,
	}, nil
}

// EmailCheck check if email exists in DB
func (r *UserRepository) CheckSocialNetworkID(social_network_id string) (*bool, error) {
	var idIsExist bool
	err := r.Store.Db.QueryRow("SELECT EXISTS (SELECT social_network_id FROM users WHERE social_network_id = $1)", social_network_id).Scan(&idIsExist)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while social_network_id checking. Err msg: %v", err)
		return &idIsExist, err
	}
	return &idIsExist, nil
}


// FindByEmail searchs and returns user by email
func (r *UserRepository) FindBySocialNetworkId(socialNetworkID string) (*model.UserDTO, error) {
	user := &model.UserDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM users WHERE social_network_id = $1",
	socialNetworkID).Scan(
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
		&user.SocialNetwork,
		&user.SocialNetworkID,
	); err != nil {
		r.Store.Logger.Errorf("Eror occure while searching user by social network id. Err msg: %v", err)
		return nil, err
	}
	return user, nil
}