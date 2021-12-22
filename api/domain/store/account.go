package store

import (
	"golang.org/x/crypto/bcrypt"
)

// Account ...
type Account struct {
	AccountID int    `json:"accountId"`
	Login     string `json:"login"`
	Password  string `json:"-"`
}

// NewAccount creates Account with encrypted password
func NewAccount(id int, login, password string) Account {
	acc := Account{id, login, password}
	acc.Password, _ = encryptPassword(acc.Password)
	return acc
}

// setPassword sets encrypted password to Account
func (a *Account) setPassword(s string) {
	a.Password, _ = encryptPassword(s)
}

// encryptPassword
func encryptPassword(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// CheckPasswordHash matches password with encrypted password<returns true/false
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
