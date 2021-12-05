package dto

// AccountDto ...
type AccountDto struct {
	AccountID int    `json:"accountId"`
	Login     string `json:"login"`
	Password  string `json:"-"`
}
