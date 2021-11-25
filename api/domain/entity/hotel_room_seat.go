package entity

// HotelRoomSeat struct
type HotelRoomSeat struct {
	HotelRoomSeatID int
	Description     string
	IsFree          bool
}

// SetDescription sets Hotel Room Seats description
func (h *HotelRoomSeat) SetDescription(s string) {
	h.Description = s
}

// SetSeatStatus sets Hotel Room Seat status (free/occupied)
func (h *HotelRoomSeat) SetSeatStatus(b bool) {
	h.IsFree = b
}
