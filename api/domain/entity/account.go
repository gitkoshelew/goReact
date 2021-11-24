package entity

import "fmt"

// Account ...
type Account struct {
	AccountID int
	Login     string
	Password  string
}

func (a *Account) getInfo() string {
	return fmt.Sprintf("Account ID: %d\n"+
		"Login: %s\n",
		a.AccountID, a.Login)
}

func (a *Account) setPassword(s string) {
	a.Password = s
}
