package entity

// Account ...
type Account struct {
	AccountID int
	Login     string
	Password  string
}

// set Accounts password
func (a *Account) setPassword(s string) {
	a.Password = s
}
