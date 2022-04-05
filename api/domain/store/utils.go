package store

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash if passwords are same err=nil
func (s *Store) CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		s.Logger.Errorf("Eror occured while checking users email or password. Err msg: %v", err)
		return err
	}
	return nil
}

// EncryptPassword ...
func (s *Store) EncryptPassword(password *string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.MinCost)
	if err != nil {
		s.Logger.Errorf("Eror occured while password encrypting. Err msg: %v", err)
		return err
	}
	*password = string(b)
	return nil
}

// StringOfArrayFromJSONToPSQL ...
func StringOfArrayFromJSONToPSQL(str *string) {
	*str = strings.Replace(*str, "[", "{", -1)
	*str = strings.Replace(*str, "]", "}", -1)
	*str = strings.Replace(*str, " ", ", ", -1)
}
