package dto

import (
	"goReact/domain/entity"
)

// GetAccountsDto ...
func GetAccountsDto() []AccountDto {
	accounts := entity.GetAccounts()
	return []AccountDto{
		AccountToDto(accounts[0]),
		AccountToDto(accounts[1]),
		AccountToDto(accounts[2]),
		AccountToDto(accounts[3]),
		AccountToDto(accounts[4]),
	}
}

// GetUsersDto ...
func GetUsersDto() []UserDto {
	users := entity.GetUsers()
	return []UserDto{
		UserToDto(users[0]),
		UserToDto(users[1]),
		UserToDto(users[2]),
		UserToDto(users[3]),
		UserToDto(users[4]),
	}
}

// GetEmployeesDto ...
func GetEmployeesDto() []EmployeeDto {
	employees := entity.GetEmployees()
	return []EmployeeDto{
		EmployeeToDto(employees[0]),
		EmployeeToDto(employees[1]),
	}
}

// GetClientsDto ...
func GetClientsDto() []ClientDto {
	clients := entity.GetClients()
	return []ClientDto{
		ClientToDto(clients[0]),
		ClientToDto(clients[1]),
		ClientToDto(clients[2]),
	}
}

// GetPetsDto ...
func GetPetsDto() []PetDto {
	pets := entity.GetPets()
	return []PetDto{
		PetToDto(pets[0]),
		PetToDto(pets[1]),
		PetToDto(pets[2]),
	}
}

// GetHotelsDto ...
func GetHotelsDto() []HotelDto {
	hotels := entity.GetHotels()
	return []HotelDto{
		HotelToDto(hotels[0]),
		HotelToDto(hotels[1]),
		HotelToDto(hotels[2]),
	}
}

// GetHotelRoomsDto ...
func GetHotelRoomsDto() []RoomDto {
	rooms := entity.GetHotelRooms()
	return []RoomDto{
		RoomToDto(rooms[0]),
		RoomToDto(rooms[1]),
		RoomToDto(rooms[2]),
		RoomToDto(rooms[3]),
	}
}

// GetHotelRoomSeatsDto ...
func GetHotelRoomSeatsDto() []SeatDto {
	seats := entity.GetHotelRoomSeats()
	return []SeatDto{
		SeatToDto(seats[0]),
		SeatToDto(seats[1]),
		SeatToDto(seats[2]),
		SeatToDto(seats[3]),
		SeatToDto(seats[4]),
		SeatToDto(seats[5]),
	}
}

// GetBookingsDto ...
func GetBookingsDto() []BookingDto {
	bookings := entity.GetBookings()
	return []BookingDto{
		BookingToDto(bookings[0]),
		BookingToDto(bookings[1]),
		BookingToDto(bookings[2]),
	}
}
