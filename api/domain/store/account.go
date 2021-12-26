package store

// Account ...
type Account struct {
	AccountID int    `json:"accountId"`
	Login     string `json:"login"`
	Password  string `json:"-"`
}
