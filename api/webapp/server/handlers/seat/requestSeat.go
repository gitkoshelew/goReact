package seat

type seatRequest struct {
	SeatID      int    `json:"seatId"`
	Description string `json:"desc"`
	IsFree      bool   `json:"isFree"`
	RoomID      int    `json:"roomId"`
}
