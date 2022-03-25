package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Login ...
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate ...
func (l *Login) Validate() error {
	return validation.ValidateStruct(
		l,
		validation.Field(&l.Email, validation.Required, is.Email, validation.By(IsSQL)),
		validation.Field(&l.Password, validation.Required, validation.Length(5, 25), validation.By(IsSQL)),
	)
}
