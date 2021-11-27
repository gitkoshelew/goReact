package server

import (
	"goReact/domain/entity"
	"goReact/pkg/date"

	"cloud.google.com/go/civil"
)

// TestStorage ...
type TestStorage struct {
	Users entity.User
}

// GetAccounts ...
func GetAccounts() []entity.Account {

	accounts := []entity.Account{
		entity.NewAccount(1, "login_1", "password_1"),
		entity.NewAccount(2, "login_2", "password_2"),
		entity.NewAccount(3, "login_3", "password_3"),
		entity.NewAccount(4, "login_4", "password_4"),
		entity.NewAccount(5, "login_5", "password_5"),
	}
	return accounts

}

// GetUsers ...
func GetUsers() []entity.User {
	accounts := GetAccounts()

	users := []entity.User{
		entity.User{accounts[0], 1, "Ivan", "Ivanov", "Ivanovich", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Lenina 10", "80252552525", "ivan@mail.com"},
		entity.User{accounts[1], 2, "Sergey", "Sergeyev", "Sergeyvich", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Krasnaya 177", "8025123131", "sergey@mail.com"},
		entity.User{accounts[2], 3, "Vladimir", "Vladimirov", "Vladimirovich", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Vulica 30", "802521231235", "vladimir@mail.com"},
		entity.User{accounts[3], 4, "Olga", "Oleinikova", "Olegovna", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Nemiga 65", "8025944655", "olga@mail.com"},
		entity.User{accounts[4], 5, "Valeria", "Valerianovna", "Valerievna", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Bobruyskaya 7", "80252545645", "valeria@mail.com"},
	}
	return users
}

// GetEmployees ...
func GetEmployees() []entity.Employee {
	users := GetUsers()
	hotels := GetHotels()
	employees := []entity.Employee{
		entity.Employee{users[0], 1, hotels[0], "Manager", "Owner"},
		entity.Employee{users[1], 2, hotels[1], "Administrator", "Employee"},
	}
	return employees
}

// GetClients ...
func GetClients() []entity.Client {
	users := GetUsers()
	pets := GetPets()
	clients := []entity.Client{
		entity.Client{users[2], 1, []entity.Pet{pets[0]}, nil},
		entity.Client{users[3], 2, []entity.Pet{pets[1]}, nil},
		entity.Client{users[4], 3, []entity.Pet{pets[2]}, nil},
	}
	return clients
}

//GetPets ...
func GetPets() []entity.Pet {
	pets := []entity.Pet{
		entity.Pet{1, "Barsik", entity.PetTypeCat, 3, 4.6, "none"},
		entity.Pet{2, "Murzik", entity.PetTypeCat, 4, 4.6, "none"},
		entity.Pet{3, "Bobik", entity.PetTypeDog, 5, 4.6, "none"},
	}
	return pets
}

// GetHotels ...
func GetHotels() []entity.Hotel {
	hotels := []entity.Hotel{
		entity.Hotel{1, "Hotel-1", "Address-1", []entity.HotelRoom{entity.HotelRoom{1, 101, entity.PetTypeCat, []entity.HotelRoomSeat{entity.HotelRoomSeat{1, "Seat 1 of room 101, hotel 1, Pet type: cat", true}}}}, nil},
		entity.Hotel{2, "Hotel-2", "Address-2", []entity.HotelRoom{entity.HotelRoom{2, 202, entity.PetTypeCat, []entity.HotelRoomSeat{entity.HotelRoomSeat{2, "Seat 2 of room 202, hotel 2, Pet type: cat", true}}}}, nil},
		entity.Hotel{3, "Hotel-3", "Address-3", []entity.HotelRoom{entity.HotelRoom{3, 303, entity.PetTypeDog, []entity.HotelRoomSeat{entity.HotelRoomSeat{3, "Seat 3 of room 303, hotel 3, Pet type: dog", false}}}}, nil},
	}
	return hotels
}

// GetHotelRooms ...
func GetHotelRooms() []entity.HotelRoom {
	seats := GetHotelRoomSeats()
	rooms := []entity.HotelRoom{
		entity.HotelRoom{1, 101, entity.PetTypeCat, []entity.HotelRoomSeat{seats[0]}},
		entity.HotelRoom{2, 202, entity.PetTypeCat, []entity.HotelRoomSeat{seats[1], seats[3]}},
		entity.HotelRoom{3, 303, entity.PetTypeDog, []entity.HotelRoomSeat{seats[2]}},
		entity.HotelRoom{4, 404, entity.PetTypeDog, []entity.HotelRoomSeat{seats[4], seats[5]}},
	}
	return rooms
}

// GetHotelRoomSeats ...
func GetHotelRoomSeats() []entity.HotelRoomSeat {
	hotelRoomSeats := []entity.HotelRoomSeat{
		entity.HotelRoomSeat{1, "Seat for cat", false},
		entity.HotelRoomSeat{2, "Seat for cat", false},
		entity.HotelRoomSeat{3, "Seat for dog", false},
		entity.HotelRoomSeat{4, "Seat for cat", true},
		entity.HotelRoomSeat{5, "Seat for dog", true},
		entity.HotelRoomSeat{6, "Seat for dog", true},
	}
	return hotelRoomSeats
}

// GetBookings ...
func GetBookings() []entity.Booking {
	pets := GetPets()
	seats := GetHotelRoomSeats()
	employees := GetEmployees()
	bookings := []entity.Booking{
		entity.Booking{1, &pets[0], &seats[0], entity.BookingStatusInProgress, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 21}}, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 30}}, employees[0], "Booking note1"},
		entity.Booking{2, &pets[1], &seats[1], entity.BookingStatusCompleted, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 19}}, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 20}}, employees[0], "Booking note2"},
		entity.Booking{3, &pets[2], &seats[2], entity.BookingStatusPending, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 30}}, date.Date{Date: civil.Date{Year: 2021, Month: 12, Day: 10}}, employees[1], "Booking note3"},
	}
	return bookings
}
