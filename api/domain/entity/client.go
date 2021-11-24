package entity

// Client extends User and has all User (and Account) fields
type Client struct {
	User
	ClientID int
	Pets     Pet
	Bookings Booking
}

// set Clients Bookings
func (c *Client) setBookings(b Booking) {
	c.Bookings = b
}

// set Clients Pets
func (c *Client) setPets(p Pet) {
	c.Pets = p
}

// Add pets to Client - not realized
func (c *Client) addPet(p Pet) {
}

// Add Booking to Client - not realized
func (c *Client) addBooking(b Booking) {
}
