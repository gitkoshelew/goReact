package client

type clientRequest struct {
	UserID      int   `json:"userId"`
	ClientID    int   `json:"clientId"`
	PetsIDs     []int `json:"petIds"`
	BookingsIDs []int `json:"bookingIds"`
}
