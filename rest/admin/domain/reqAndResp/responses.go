package reqandresp

import "time"

// SeatsAndDates ...
type SeatsAndDates struct {
	SeatID    int       `json:"seatId"`
	StratDate time.Time `json:"start"`
	EndDate   time.Time `json:"end"`
}
