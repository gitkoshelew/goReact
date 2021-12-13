package store

// Hotel struct
type Hotel struct {
	HotelID int    `json:"hotelId"`
	Name    string `json:"nameId"`
	Address string `json:"addressId"`
}

// SetName sets Hotels Name
func (h *Hotel) SetName(s string) {
	h.Name = s
}

// SetAddress sets Hotels Address
func (h *Hotel) SetAddress(s string) {
	h.Address = s
}
