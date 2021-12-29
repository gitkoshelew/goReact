package dto

import (
	"goReact/domain/store"
)

// AccountToDto makes a DTO from Account object
func AccountToDto(a store.Account) AccountDto {
	return AccountDto{
		AccountID: a.AccountID,
		Login:     a.Login,
		Password:  a.Password,
	}

}

// UserToDto makes DTO from user object
func UserToDto(u store.User) UserDto {
	return UserDto{
		UserID:      u.UserID,
		Email:       u.Email,
		Password:    u.Password,
		Role:        string(u.Role),
		Verified:    u.Verified,
		Name:        u.Name,
		Surname:     u.Surname,
		MiddleName:  u.MiddleName,
		Sex:         string(u.Sex),
		DateOfBirth: u.DateOfBirth,
		Address:     u.Address,
		Phone:       u.Phone,
		Photo:       u.Photo,
	}
}

// EmployeeToDto makes DTO from Employee object
func EmployeeToDto(e store.Employee) EmployeeDto {
	return EmployeeDto{
		EmployeeID: e.EmployeeID,
		UserID:     e.UserID,
		HotelID:    e.Hotel.HotelID,
		Position:   e.Position,
		Role:       e.Role,
	}
}

// ClientToDto makes DTO from Client object
func ClientToDto(c store.Client) ClientDto {
	return ClientDto{
		ClientID: c.ClientID,
		UserID:   c.UserID}
}

// PetToDto makes DTO from Pet object
func PetToDto(p store.Pet) PetDto {
	return PetDto{
		PetID:     p.PetID,
		Name:      p.Name,
		Type:      string(p.Type),
		Weight:    p.Weight,
		Diesieses: p.Diesieses,
		OwnerID:   p.Owner.ClientID}
}

// SeatToDto makes DTO from HotelRoomSeat object
func SeatToDto(s store.Seat) SeatDto {
	return SeatDto{
		SeatID:      s.SeatID,
		Description: s.Description,
		IsFree:      s.IsFree,
		RoomID:      s.Room.RoomID,
	}
}

// RoomToDto makes DTO from HotelRoom object
func RoomToDto(r store.Room) RoomDto {
	return RoomDto{
		RoomID:     r.RoomID,
		RoomNumber: r.RoomNumber,
		PetType:    string(r.PetType),
		HotelID:    r.Hotel.HotelID}
}

// HotelToDto makes DTO from Hotel object
func HotelToDto(h store.Hotel) HotelDto {
	return HotelDto{
		HotelID: h.HotelID,
		Name:    h.Name,
		Address: h.Address,
	}
}

// BookingToDto makes DTO from Booking object
func BookingToDto(b store.Booking) BookingDto {
	return BookingDto{
		BookingID:   b.BookingID,
		SeatID:      b.Seat.SeatID,
		PetID:       b.Pet.PetID,
		EmployeeID:  b.Employee.EmployeeID,
		Status:      string(b.Status),
		StartDate:   b.StartDate,
		EndDate:     b.EndDate,
		ClientNotes: b.ClientNotes,
	}
}
