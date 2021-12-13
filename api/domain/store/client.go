package store

// Client extends User and has all User (and Account) fields
type Client struct {
	ClientID int `json:"clientId"`
	User
}
