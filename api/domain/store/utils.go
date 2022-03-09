package store

import "golang.org/x/crypto/bcrypt"

// CheckPasswordHash if passwords are same err=nil
func (s *Store) CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	s.Logger.Errorf("Eror occured while checking users email or password. Err msg: %w", err)
	return err
}

// EncryptPassword ...
func (s *Store) EncryptPassword(passwod *string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(*passwod), bcrypt.MinCost)
	if err != nil {
		s.Logger.Errorf("Eror occured while password encrypting. Err msg: %w", err)
		return err
	}
	*passwod = string(b)
	return nil
}
