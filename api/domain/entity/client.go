package entity

// Client extends User and has all User (and Account) fields
type Client struct {
	User
	ClientID int
	Pets     []Pet
	Bookings []Booking
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
