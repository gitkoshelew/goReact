package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	UserID      int        `json:"userId" csv:"userId"`
	Email       string     `json:"email" csv:"email"`
	Password    string     `json:"-" csv:"-"`
	Role        Role       `json:"role" csv:"role"`
	Verified    *bool      `json:"verified" csv:"verified"`
	Name        string     `json:"name" csv:"name"`
	Surname     string     `json:"sName" csv:"sName"`
	MiddleName  string     `json:"mName" csv:"mName"`
	Sex         Sex        `json:"sex" csv:"sex"`
	DateOfBirth *time.Time `json:"birthDate" csv:"birthDate"`
	Address     string     `json:"address" csv:"address"`
	Phone       string     `json:"phone" csv:"phone"`
	Photo       string     `json:"photo" csv:"photo"`
}

// UserDTO ...
type UserDTO struct {
	UserID      int        `json:"userId" csv:"userId"`
	Email       string     `json:"email" csv:"email"`
	Password    string     `json:"password" csv:"-"`
	Role        string     `json:"role,omitempty" csv:"role"`
	Verified    *bool      `json:"verified,omitempty" csv:"verified"`
	Name        string     `json:"name" csv:"name"`
	Surname     string     `json:"sName" csv:"sName"`
	MiddleName  string     `json:"mName" csv:"mName"`
	Sex         string     `json:"sex" csv:"sex"`
	DateOfBirth *time.Time `json:"birthDate" csv:"birthDate"`
	Address     string     `json:"address" csv:"address"`
	Phone       string     `json:"phone" csv:"phone"`
	Photo       string     `json:"photo" csv:"photo"`
}

// Role ...
type Role string

// Role constants
const (
	ClientRole    Role = "client"
	EmployeeRole  Role = "employee"
	AnonymousRole Role = "anonymous"
)

// Sex ...
type Sex string

// Sex constants
const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

// Validate ...
func (u *UserDTO) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email, validation.By(IsSQL)),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 100), validation.By(IsSQL)),
		validation.Field(&u.Role, validation.By(IsRole)),
		validation.Field(&u.Name, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(2, 30), validation.By(IsSQL)),
		validation.Field(&u.Surname, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(2, 30), validation.By(IsSQL)),
		validation.Field(&u.MiddleName, validation.By(IsLetterHyphenSpaces), validation.Length(0, 30), validation.By(IsSQL)),
		validation.Field(&u.Sex, validation.Required, validation.By(IsSex)),
		validation.Field(&u.DateOfBirth, validation.Required, validation.By(IsValidBirthDate)),
		validation.Field(&u.Address, validation.Required, validation.By(IsSQL), validation.Length(10, 40)),
		validation.Field(&u.Phone, validation.Required, validation.By(IsPhone)),
		validation.Field(&u.Photo, validation.By(IsSQL)),
	)
}

// EncryptPassword ...
func EncryptPassword(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// CheckPasswordHash if passwords are same err=nil
func CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
