package account

type accountRequest struct {
	AccountID int    `json:"accountId"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}
