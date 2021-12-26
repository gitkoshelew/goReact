package store

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Role ...
type Role string

// Role constants
const (
	ClientRole    Role = "client"
	EmployeeRole  Role = "employee"
	AnonymousRole Role = "anonymous"
)

// User extends Account and has all Account fields
type User struct {
	UserID      int       `json:"userId"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Role        Role      `json:"role"`
	Verified    bool      `json:"verified"`
	Name        string    `json:"name"`
	Surname     string    `json:"sName"`
	MiddleName  string    `json:"mName"`
	Sex         int       `json:"sex"`
	DateOfBirth time.Time `json:"birthDate"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Photo       string    `json:"photo"`
}

// NewUser creates User with encrypted password
func NewUser(id, sex int, email, password, role, name, surname, middleName, address, phone, photo string, verified bool, dateOfBirth time.Time) User {
	user := User{
		UserID:      id,
		Email:       email,
		Password:    password,
		Role:        Role(role),
		Verified:    verified,
		Name:        name,
		Surname:     surname,
		MiddleName:  middleName,
		Sex:         sex,
		DateOfBirth: dateOfBirth,
		Address:     address,
		Phone:       phone,
		Photo:       photo,
	}
	user.Password, _ = EncryptPassword(user.Password)
	return user
}

// EncryptPassword ...
func EncryptPassword(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// CheckPasswordHash matches password with encrypted password<returns true/false
func CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
