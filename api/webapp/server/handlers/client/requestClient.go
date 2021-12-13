package client

type clientRequest struct {
	ClientID int `json:"clientId"`
	UserID   int `json:"userId"`
}
