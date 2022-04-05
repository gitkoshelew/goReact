package store

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// StringOfArrayFromJSONToPSQL ...
func StringOfArrayFromJSONToPSQL(str *string) {
	*str = strings.Replace(*str, "[", "{", -1)
	*str = strings.Replace(*str, "]", "}", -1)
	*str = strings.Replace(*str, " ", ", ", -1)
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
