package entity

// HotelRoomSeat struct
type HotelRoomSeat struct {
	HotelRoomSeatID int    `json:"seatId"`
	Description     string `json:"desc"`
	IsFree          bool   `json:"isFree"`
}

// SetDescription sets Hotel Room Seats description
func (h *HotelRoomSeat) SetDescription(s string) {
	h.Description = s
}

// SetSeatStatus sets Hotel Room Seat status (free/occupied)
func (h *HotelRoomSeat) SetSeatStatus(b bool) {
	h.IsFree = b
}

// GetSeatByID returns Seat by id from storage
func GetSeatByID(id int) HotelRoomSeat {
	var seat HotelRoomSeat
	for _, s := range GetHotelRoomSeats() {
		if id == s.HotelRoomSeatID {
			seat = s
		}
	}
	return seat
}

// GetSeatsByID returns []HotelRoomSeat by []ids from storage
func GetSeatsByID(ids []int) []HotelRoomSeat {
	var seats []HotelRoomSeat
	for _, id := range ids {
		seats = append(seats, GetSeatByID(id))
	}
	return seats
}
