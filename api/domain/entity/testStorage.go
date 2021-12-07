package entity

import (
	"goReact/pkg/date"

	"cloud.google.com/go/civil"
)

// GetAccounts ...
func GetAccounts() []Account {
	return []Account{
		NewAccount(1, "login_1", "password_1"),
		NewAccount(2, "login_2", "password_2"),
		NewAccount(3, "login_3", "password_3"),
		NewAccount(4, "login_4", "password_4"),
		NewAccount(5, "login_5", "password_5"),
	}
}

// GetUsers ...
func GetUsers() []User {
	accounts := GetAccounts()
	return []User{
		User{accounts[0], 1, "Ivan", "Ivanov", "Ivanovich", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Lenina 10", "80252552525", "ivan@mail.com"},
		User{accounts[1], 2, "Sergey", "Sergeyev", "Sergeyvich", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Krasnaya 177", "8025123131", "sergey@mail.com"},
		User{accounts[2], 3, "Vladimir", "Vladimirov", "Vladimirovich", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Vulica 30", "802521231235", "vladimir@mail.com"},
		User{accounts[3], 4, "Olga", "Oleinikova", "Olegovna", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Nemiga 65", "8025944655", "olga@mail.com"},
		User{accounts[4], 5, "Valeria", "Valerianovna", "Valerievna", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Bobruyskaya 7", "80252545645", "valeria@mail.com"},
	}
}

// GetEmployees ...
func GetEmployees() []Employee {
	users := GetUsers()
	hotels := GetHotels()
	return []Employee{
		Employee{users[0], hotels[0], 1, "Manager", "Owner"},
		Employee{users[1], hotels[1], 2, "Administrator", "Employee"},
	}
}

// GetClients ...
func GetClients() []Client {
	users := GetUsers()
	pets := GetPets()
	return []Client{
		Client{users[2], 1, []Pet{pets[0]}, nil},
		Client{users[3], 2, []Pet{pets[1]}, nil},
		Client{users[4], 3, []Pet{pets[2]}, nil},
	}
}

//GetPets ...
func GetPets() []Pet {
	return []Pet{
		Pet{1, "Barsik", PetTypeCat, 3, 4.6, "none"},
		Pet{2, "Murzik", PetTypeCat, 4, 4.6, "none"},
		Pet{3, "Bobik", PetTypeDog, 5, 4.6, "none"},
	}
}

// GetHotels ...
func GetHotels() []Hotel {
	return []Hotel{
		Hotel{1, "Hotel-1", "Address-1", []HotelRoom{HotelRoom{1, 101, PetTypeCat, []HotelRoomSeat{HotelRoomSeat{1, "Seat 1 of room 101, hotel 1, Pet type: cat", true}}}}, nil},
		Hotel{2, "Hotel-2", "Address-2", []HotelRoom{HotelRoom{2, 202, PetTypeCat, []HotelRoomSeat{HotelRoomSeat{2, "Seat 2 of room 202, hotel 2, Pet type: cat", true}}}}, nil},
		Hotel{3, "Hotel-3", "Address-3", []HotelRoom{HotelRoom{3, 303, PetTypeDog, []HotelRoomSeat{HotelRoomSeat{3, "Seat 3 of room 303, hotel 3, Pet type: dog", false}}}}, nil},
	}
}

// GetHotelRooms ...
func GetHotelRooms() []HotelRoom {
	seats := GetHotelRoomSeats()
	return []HotelRoom{
		HotelRoom{1, 101, PetTypeCat, []HotelRoomSeat{seats[0]}},
		HotelRoom{2, 202, PetTypeCat, []HotelRoomSeat{seats[1], seats[3]}},
		HotelRoom{3, 303, PetTypeDog, []HotelRoomSeat{seats[2]}},
		HotelRoom{4, 404, PetTypeDog, []HotelRoomSeat{seats[4], seats[5]}},
	}
}

// GetHotelRoomSeats ...
func GetHotelRoomSeats() []HotelRoomSeat {
	return []HotelRoomSeat{
		HotelRoomSeat{1, "Seat for cat", false},
		HotelRoomSeat{2, "Seat for cat", false},
		HotelRoomSeat{3, "Seat for dog", false},
		HotelRoomSeat{4, "Seat for cat", true},
		HotelRoomSeat{5, "Seat for dog", true},
		HotelRoomSeat{6, "Seat for dog", true},
	}
}

// GetBookings ...
func GetBookings() []Booking {
	pets := GetPets()
	seats := GetHotelRoomSeats()
	employees := GetEmployees()
	return []Booking{
		Booking{1, pets[0], seats[0], BookingStatusInProgress, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 21}}, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 30}}, employees[0], "Booking note1"},
		Booking{2, pets[1], seats[1], BookingStatusCompleted, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 19}}, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 20}}, employees[0], "Booking note2"},
		Booking{3, pets[2], seats[2], BookingStatusPending, date.Date{Date: civil.Date{Year: 2021, Month: 11, Day: 30}}, date.Date{Date: civil.Date{Year: 2021, Month: 12, Day: 10}}, employees[1], "Booking note3"},
	}
}
