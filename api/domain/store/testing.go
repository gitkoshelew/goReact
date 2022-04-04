package store

import (
	"goReact/domain/model"
	"goReact/webapp"
	"strconv"
	"testing"
)

// ID ...
type ID struct {
	User     int
	Employee int
	Pet      int
	Hotel    int
	Room     int
	Seat     int
	Booking  int
	Image    int
}

// TestStore ...
func TestStore(t *testing.T, host, dbName, user, password, port, sslMode string) (*Store, func()) {
	t.Helper()

	config := &webapp.Config{}
	config.NewConfig()
	config.DbConnection.Host = host
	config.DbConnection.DbName = dbName
	config.DbConnection.Username = user
	config.DbConnection.Password = password
	config.DbConnection.Port, _ = strconv.Atoi(port)
	config.DbConnection.Sslmode = sslMode

	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func() {
		s.Open()
		_, err := s.Db.Exec("TRUNCATE users, hotel, employee, room, seat, pet, booking CASCADE")
		if err != nil {
			t.Fatal(err)
		}
		s.Close()
	}
}

// FillDB ...
func FillDB(t *testing.T, s *Store) *ID {
	s.Open()

	user := model.TestUser()
	user.Email = "test@mail.org"
	userID, _ := s.User().Create(user)
	user.UserID = *userID

	hotel := model.TestHotel()
	hotelID, _ := s.Hotel().Create(hotel)
	hotel.HotelID = *hotelID

	employee := model.TestEmployee()
	employee.User = *user
	employee.Hotel = *hotel
	employeeID, _ := s.Employee().Create(employee)
	employee.EmployeeID = *employeeID

	room := model.TestRoom()
	room.Hotel = *hotel
	roomID, _ := s.Room().Create(room)
	room.RoomID = *roomID

	seat := model.TestSeat()
	seat.Room = *room
	seatID, _ := s.Seat().Create(seat)
	seat.SeatID = *seatID

	pet := model.TestPet()
	pet.Owner = *user
	petID, _ := s.Pet().Create(pet)
	pet.PetID = *petID

	booking := model.TestBooking()
	booking.Pet = *pet
	booking.Seat = *seat
	booking.Employee = *employee
	bookingID, _ := s.Booking().Create(booking)
	booking.BookingID = *bookingID

	image := model.TestImage()
	image.OwnerID = user.UserID
	imageID, _ := s.Image().Create(image)
	image.ImageID = *imageID

	s.Close()

	return &ID{
		User:     *userID,
		Employee: *employeeID,
		Pet:      *petID,
		Hotel:    *hotelID,
		Room:     *roomID,
		Seat:     *seatID,
		Booking:  *bookingID,
		Image:    *imageID,
	}
}
