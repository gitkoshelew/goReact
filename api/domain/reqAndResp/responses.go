package reqandresp

import "time"

// RoomInfo ...
type RoomInfo struct {
	SeatID    int        `json:"seatId"`
	StartDate *time.Time `json:"start"`
	EndDate   *time.Time `json:"end"`
}
