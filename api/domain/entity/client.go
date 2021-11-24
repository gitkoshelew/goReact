package entity

import "fmt"

// Client ...
type Client struct {
	User
	ClientID int
	Pets     string
	Bookings string
}

func (c *Client) getInfo() string {
	return fmt.Sprintf("Account ID: %d\n"+
		"Login: %s\n"+
		"User ID: %d\n"+
		"Name: %s\n"+
		"Surname: %s\n"+
		"Middlename: %s\n"+
		"Date of birth: %s\n"+
		"Address: %s\n"+
		"Phone number: %s\n"+
		"Email: %s\n"+
		"Client ID: %d\n"+
		"Pets: %s=\n"+
		"Bookings: %s\n",
		c.AccountID, c.Login, c.UserID, c.Name, c.Surname, c.MiddleName, c.DateOfBirth, c.Address, c.Phone, c.Email, c.ClientID, c.Pets, c.Bookings)

	// %v  ||  "data: %v, c"
}

func (c *Client) setBookings(b string) {
	c.Bookings = b
}

func (c *Client) setPets(p string) {
	c.Pets = p
}

// addPet
func (c *Client) addPet(p string) {

}

// addBooking
func (c *Client) addBooking(b string) {

}
