package model

import (
	"authentication/pkg/logging"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	UserID   int    `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	Verified bool   `json:"verified"`
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

var (
	logger = logging.GetLogger()
)

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 100)),
		validation.Field(&u.Role, validation.Required, validation.By(IsRole)),
	)
}

// WithEncryptedPassword creates User with encrypted password
func (u *User) WithEncryptedPassword() error {
	EncryptedPassword, err := EncryptPassword(u.Password)
	if err != nil {
		logger.Errorf("bad request. err msg:%v.", err)
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
