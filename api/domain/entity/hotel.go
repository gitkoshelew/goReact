package entity

import "errors"

// Hotel struct
type Hotel struct {
	HotelID  int         `json:"hotelId"`
	Name     string      `json:"nameId"`
	Address  string      `json:"addressId"`
	Rooms    []HotelRoom `json:"roomsIds"`
	Bookings []Booking   `json:"bookingsIds"`
}

// SetName sets Hotels Name
func (h *Hotel) SetName(s string) {
	h.Name = s
}

// SetAddress sets Hotels Address
func (h *Hotel) SetAddress(s string) {
	h.Address = s
}

// SetRooms sets Hotels Rooms
func (h *Hotel) SetRooms(r []HotelRoom) {
	h.Rooms = r
}

// SetBookings sets Hotels Bookings
func (h *Hotel) SetBookings(b []Booking) {
	h.Bookings = b
}

//RoomsCount gets number of Hotels Rooms
func (h *Hotel) RoomsCount() int {
	if h == nil {
		return 0
	}
	return len(h.Rooms)
}

// SeatsCount gets number of all Hotel seats
func (h *Hotel) SeatsCount() int {
	count := 0
	for _, room := range h.Rooms {
		count += room.SeatsCount()
	}
	return count
}

// FreeSeatsCount returns number of freee seats in hotel during pet type. Example: "entity.Hotel.FreeSeatsCount(PetTypeCat)"
// If entered wrong pet type - recieve 0 and error
func (h *Hotel) FreeSeatsCount(p PetType) (int, error) {
	count := 0

	switch p {

	case PetTypeCat:
		for _, room := range h.Rooms {
			if room.PetType == PetTypeCat {
				for _, seat := range room.Seats {
					if seat.IsFree {
						count++
					}
				}
			}
		}
		return count, nil

	case PetTypeDog:
		for _, room := range h.Rooms {
			if room.PetType == PetTypeDog {
				for _, seat := range room.Seats {
					if seat.IsFree {
						count++
					}
				}
			}
		}
		return count, nil
	default:
		err := errors.New("Incorrect pet type. Please check the entered data and try again! ")
		return count, err
	}
}

// GetHotelByID returns Hotel by id from storage
func GetHotelByID(id int) Hotel {
	hotels := GetHotels()
	var hotel Hotel
	for _, h := range hotels {
		if id == h.HotelID {
			hotel = h
		}
	}
	return hotel
}
