package dto

// ClientDto ...
type ClientDto struct {
	UserID     int   `json:"userId"`
	ClientID   int   `json:"clientId"`
	PetsID     []int `json:"petIds"`
	BookingsID []int `json:"bookingIds"`
}
