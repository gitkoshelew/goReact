package dto

import "goReact/domain/entity"

// AccountToDto makes a DTO from Account object
func AccountToDto(a entity.Account) AccountDto {
	return AccountDto{
		AccountID: a.AccountID,
		Login:     a.Login,
		Password:  a.Password,
	}
}

// UserToDto makes DTO from user object
func UserToDto(u entity.User) UserDto {
	return UserDto{
		AccountID:   u.AccountID,
		UserID:      u.UserID,
		Name:        u.Name,
		Surname:     u.Surname,
		MiddleName:  u.MiddleName,
		DateOfBirth: u.DateOfBirth,
		Address:     u.Address,
		Phone:       u.Phone,
		Email:       u.Email,
	}
}

// EmployeeToDto makes DTO from Employee object
func EmployeeToDto(e entity.Employee) EmployeeDto {
	return EmployeeDto{
		HotelID:    e.Hotel.HotelID,
		EmployeeID: e.EmployeeID,
		Position:   e.Position,
		Role:       e.Role,
	}
}

// ClientToDto makes DTO from Client object
func ClientToDto(c entity.Client) ClientDto {
	var petsIds []int
	for _, v := range c.Pets {
		petsIds = append(petsIds, v.PetID)
	}
	var bookingIds []int
	for _, v := range c.Bookings {
		bookingIds = append(bookingIds, v.BookingID)
	}
	return ClientDto{
		UserID:     c.UserID,
		ClientID:   c.ClientID,
		PetsID:     petsIds,
		BookingsID: bookingIds,
	}
}

// PetToDto makes DTO from Pet object
func PetToDto(p entity.Pet) PetDto {
	return PetDto{
		PetID:     p.PetID,
		Name:      p.Name,
		Type:      string(p.Type),
		OwnerID:   p.OwnerID,
		Weight:    p.Weight,
		Diesieses: p.Diesieses,
	}
}

// SeatToDto makes DTO from HotelRoomSeat object
func SeatToDto(s entity.HotelRoomSeat) SeatDto {
	return SeatDto{
		HotelRoomSeatID: s.HotelRoomSeatID,
		Description:     s.Description,
		IsFree:          s.IsFree,
	}
}

// RoomToDto makes DTO from HotelRoom object
func RoomToDto(r entity.HotelRoom) RoomDto {
	var seatsID []int
	for _, v := range r.Seats {
		seatsID = append(seatsID, v.HotelRoomSeatID)
	}
	return RoomDto{
		HotelRoomID: r.HotelRoomID,
		RoomNumber:  r.RoomNumber,
		PetType:     string(r.PetType),
		SeatsID:     seatsID,
	}
}

// HotelToDto makes DTO from Hotel object
func HotelToDto(h entity.Hotel) HotelDto {
	var bookingIds []int
	for _, v := range h.Bookings {
		bookingIds = append(bookingIds, v.BookingID)
	}
	var roomsIds []int
	for _, v := range h.Rooms {
		roomsIds = append(roomsIds, v.HotelRoomID)
	}
	return HotelDto{
		HotelID:    h.HotelID,
		Name:       h.Name,
		Address:    h.Address,
		RoomsID:    roomsIds,
		BookingsID: bookingIds,
	}
}

// BookingToDto makes DTO from Booking object
func BookingToDto(b entity.Booking) BookingDto {
	return BookingDto{
		BookingID:   b.BookingID,
		PetID:       b.Pet.PetID,
		SeatID:      b.Seat.HotelRoomSeatID,
		Status:      string(b.Status),
		StartDate:   b.StartDate,
		EndDate:     b.EndDate,
		EmployeeID:  b.Employee.EmployeeID,
		ClientNotes: b.ClientNotes,
	}
}
