package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
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
	Sex         Sex       `json:"sex"`
	DateOfBirth time.Time `json:"birthDate"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Photo       string    `json:"photo"`
}

// UserDTO ...
type UserDTO struct {
	UserID      int       `json:"userId"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	Verified    bool      `json:"verified"`
	Name        string    `json:"name"`
	Surname     string    `json:"sName"`
	MiddleName  string    `json:"mName"`
	Sex         string    `json:"sex"`
	DateOfBirth time.Time `json:"birthDate"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Photo       string    `json:"photo"`
}

// Role ...
type Role string

// Role constants
const (
	ClientRole    Role = "client"
	EmployeeRole  Role = "employee"
	AnonymousRole Role = "anonymous"
)

// StrToRole ...
func StrToRole(s *string) Role {
	switch *s {
	case "client":
		return ClientRole
	case "employee":
		return EmployeeRole
	case "anonymous":
		return AnonymousRole
	default:
		return ""
	}
}

// Sex ...
type Sex string

// Sex constants
const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

// StrToSex ...
func StrToSex(s *string) Sex {
	switch *s {
	case "male":
		return SexMale
	case "female":
		return SexFemale
	default:
		return ""
	}
}

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 100)),
		validation.Field(&u.Role, validation.Required, validation.By(IsRole)),
		validation.Field(&u.Name, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(3, 20)),
		validation.Field(&u.Surname, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(3, 20)),
		validation.Field(&u.MiddleName, validation.By(IsLetterHyphenSpaces), validation.Length(3, 20)),
		validation.Field(&u.Sex, validation.Required, validation.By(IsSex)),
		validation.Field(&u.DateOfBirth, validation.Required),
		validation.Field(&u.Address, validation.Required, validation.Length(10, 40)),
		validation.Field(&u.Phone, validation.Required, validation.By(IsPhone)),
	)
}

// WithEncryptedPassword creates User with encrypted password
func (u *User) WithEncryptedPassword() error {
	EncryptedPassword, err := EncryptPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = EncryptedPassword
	return nil
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

// ModelFromDTO ...
func (uDTO *UserDTO) ModelFromDTO() *User {
	return &User{
		UserID:      uDTO.UserID,
		Email:       uDTO.Email,
		Password:    uDTO.Password,
		Role:        StrToRole(&uDTO.Role),
		Verified:    uDTO.Verified,
		Name:        uDTO.Name,
		Surname:     uDTO.Surname,
		MiddleName:  uDTO.MiddleName,
		Sex:         StrToSex(&uDTO.Sex),
		DateOfBirth: uDTO.DateOfBirth,
		Address:     uDTO.Address,
		Phone:       uDTO.Phone,
		Photo:       uDTO.Phone,
	}
}
