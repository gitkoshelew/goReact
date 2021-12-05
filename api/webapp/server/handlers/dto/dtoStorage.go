package dto

import (
	"goReact/domain/entity"
)

// GetAccountsDto ...
func GetAccountsDto() []AccountDto {
	accounts := entity.GetAccounts()
	return []AccountDto{
		AccountDto(entity.AccountToDto(accounts[0])),
		AccountDto(entity.AccountToDto(accounts[1])),
		AccountDto(entity.AccountToDto(accounts[2])),
		AccountDto(entity.AccountToDto(accounts[3])),
		AccountDto(entity.AccountToDto(accounts[4])),
	}
}

// GetUsersDto ...
func GetUsersDto() []UserDto {
	users := entity.GetUsers()
	return []UserDto{
		UserDto(entity.UserToDto(users[0])),
		UserDto(entity.UserToDto(users[1])),
		UserDto(entity.UserToDto(users[2])),
		UserDto(entity.UserToDto(users[3])),
		UserDto(entity.UserToDto(users[4])),
	}
}

// GetEmployeesDto ...
func GetEmployeesDto() []EmployeeDto {
	employees := entity.GetEmployees()
	return []EmployeeDto{
		EmployeeDto(entity.EmployeeToDto(employees[0])),
		EmployeeDto(entity.EmployeeToDto(employees[1])),
	}
}

// GetClientsDto ...
func GetClientsDto() []ClientDto {
	clients := entity.GetClients()
	return []ClientDto{
		ClientDto(entity.ClientToDto(clients[0])),
		ClientDto(entity.ClientToDto(clients[1])),
		ClientDto(entity.ClientToDto(clients[2])),
	}
}

// GetPetsDto ...
func GetPetsDto() []PetDto {
	pets := entity.GetPets()
	return []PetDto{
		PetDto(entity.PetToDto(pets[0])),
		PetDto(entity.PetToDto(pets[1])),
		PetDto(entity.PetToDto(pets[2])),
	}
}

// GetHotelsDto ...
func GetHotelsDto() []HotelDto {
	hotels := entity.GetHotels()
	return []HotelDto{
		HotelDto(entity.HotelToDto(hotels[0])),
		HotelDto(entity.HotelToDto(hotels[1])),
		HotelDto(entity.HotelToDto(hotels[2])),
	}
}

// GetHotelRoomsDto ...
func GetHotelRoomsDto() []RoomDto {
	rooms := entity.GetHotelRooms()
	return []RoomDto{
		RoomDto(entity.RoomToDto(rooms[0])),
		RoomDto(entity.RoomToDto(rooms[1])),
		RoomDto(entity.RoomToDto(rooms[2])),
		RoomDto(entity.RoomToDto(rooms[3])),
	}
}

// GetHotelRoomSeatsDto ...
func GetHotelRoomSeatsDto() []SeatDto {
	seats := entity.GetHotelRoomSeats()
	return []SeatDto{
		SeatDto(entity.SeatToDto(seats[0])),
		SeatDto(entity.SeatToDto(seats[1])),
		SeatDto(entity.SeatToDto(seats[2])),
		SeatDto(entity.SeatToDto(seats[3])),
		SeatDto(entity.SeatToDto(seats[4])),
		SeatDto(entity.SeatToDto(seats[5])),
	}
}

// GetBookingsDto ...
func GetBookingsDto() []BookingDto {
	bookings := entity.GetBookings()
	return []BookingDto{
		BookingDto(entity.BookingToDto(bookings[0])),
		BookingDto(entity.BookingToDto(bookings[1])),
		BookingDto(entity.BookingToDto(bookings[2])),
	}
}
