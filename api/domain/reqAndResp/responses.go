package reqandresp

import "time"

// SeatsAndDates ...
type SeatsAndDates struct {
	SeatID    int       `json:"seatId"`
	StratDate time.Time `json:"start"`
	EndDate   time.Time `json:"end"`
}

//Top rooms response
type TopRooms struct {
	RoomID        int `json:"roomId"`
	TotalBookings int `json:"total bookings"`
}
