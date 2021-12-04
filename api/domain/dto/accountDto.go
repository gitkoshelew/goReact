package dto

// Account DTO
type Account struct {
	AccountID int    `json:"accountId"`
	Login     string `json:"login"`
	Password  string `json:"-"`
}
