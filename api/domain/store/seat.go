package store

// Seat struct
type Seat struct {
	SeatID      int    `json:"seatId"`
	Description string `json:"desc"`
	IsFree      bool   `json:"isFree"`
	Room        Room
}

// SetDescription sets Hotel Room Seats description
func (s *Seat) SetDescription(desc string) {
	s.Description = desc
}

// SetSeatStatus sets Hotel Room Seat status (free/occupied)
func (s *Seat) SetSeatStatus(b bool) {
	s.IsFree = b
}
