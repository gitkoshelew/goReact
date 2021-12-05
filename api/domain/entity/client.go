package entity

// Client extends User and has all User (and Account) fields
type Client struct {
	User
	ClientID int       `json:"clientId"`
	Pets     []Pet     `json:"petIds"`
	Bookings []Booking `json:"bookingIds"`
}

// SetBookings sets Clients Bookings
func (c *Client) SetBookings(b []Booking) {
	c.Bookings = b
}

// SetPets sets Clients Pets
func (c *Client) SetPets(p []Pet) {
	c.Pets = p
}

// AddPet adds pets to Client
func (c *Client) AddPet(p Pet) {
	c.Pets = append(c.Pets, p)
}

// AddBooking adds Booking to Client
func (c *Client) AddBooking(b Booking) {
	c.Bookings = append(c.Bookings, b)
}

// GetClientByID returns CLient by id from storage
func GetClientByID(id int) Client {
	clients := GetClients()
	var client Client
	for _, c := range clients {
		if id == c.ClientID {
			client = c
		}
	}
	return client
}
