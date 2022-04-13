package reqandresp

import "time"

// RoomInfo ...
type RoomInfo struct {
	SeatID    int        `json:"seatId"`
	StartDate *time.Time `json:"start"`
	EndDate   *time.Time `json:"end"`
}

// FreeSeatsResponse struct
type FreeSeatsResponse struct {
	Day        int   `json:"day"`
	SeatIDs    []int `json:"seatIds"`
	TotalCount int   `json:"totalCount"`
}

//Top rooms response
type TopRooms struct {
	RoomID        int `json:"roomId"`
	TotalBookings int `json:"total bookings"`
}
