package server

import (
	employeehandlers "hotel/internal/handlers/employeeHandlers"
	hotelhandlers "hotel/internal/handlers/hotelHandlers"
	roomhandlers "hotel/internal/handlers/roomHandlers"
	seathandlers "hotel/internal/handlers/seatHandlers"
	"hotel/internal/store"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {

	s.Router.Handle("GET", "/hotels", hotelhandlers.AllHotelsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/hotel/:id", hotelhandlers.GetHotelByID(store.New(s.Config)))
	s.Router.Handle("DELETE", "/hotel/:id", hotelhandlers.DeleteHotels(store.New(s.Config)))
	s.Router.Handle("POST", "/hotel", hotelhandlers.NewHotel(store.New(s.Config)))
	s.Router.Handle("PUT", "/hotel", hotelhandlers.UpdateHotel(store.New(s.Config)))

	s.Router.Handle("GET", "/rooms", roomhandlers.AllRoomsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/room/:id", roomhandlers.GetRoomByID(store.New(s.Config)))
	s.Router.Handle("DELETE", "/room/:id", roomhandlers.DeleteRooms(store.New(s.Config)))
	s.Router.Handle("POST", "/room", roomhandlers.NewRoom(store.New(s.Config)))
	s.Router.Handle("PUT", "/room", roomhandlers.UpdateRoom(store.New(s.Config)))

	s.Router.Handle("GET", "/seats", seathandlers.AllSeatsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/seat/:id", seathandlers.GetSeatByID(store.New(s.Config)))
	s.Router.Handle("DELETE", "/seat/:id", seathandlers.DeleteSeats(store.New(s.Config)))
	s.Router.Handle("POST", "/seat", seathandlers.NewSeat(store.New(s.Config)))
	s.Router.Handle("PUT", "/seat", seathandlers.UpdateSeat(store.New(s.Config)))

	s.Router.Handle("GET", "/employees", employeehandlers.AllEmployeeHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/employee/:id", employeehandlers.GetEmployeeByID(store.New(s.Config)))
	s.Router.Handle("DELETE", "/employee/:id", employeehandlers.DeleteEmployee(store.New(s.Config)))
	s.Router.Handle("POST", "/employee", employeehandlers.NewEmployee(store.New(s.Config)))
	s.Router.Handle("PUT", "/employee", employeehandlers.UpdateEmployee(store.New(s.Config)))

}
