package server

import (
	"goReact/domain/store"
	"goReact/webapp/admin/handlersAdmin/auth"
	bookinghandlers "goReact/webapp/admin/handlersAdmin/bookingHandlers"
	employeehandlers "goReact/webapp/admin/handlersAdmin/employeeHandlers"
	hotelhandlers "goReact/webapp/admin/handlersAdmin/hotelHandlers"
	"goReact/webapp/admin/handlersAdmin/imagehandlers"
	"goReact/webapp/admin/handlersAdmin/permission"
	pethandlers "goReact/webapp/admin/handlersAdmin/petHandlers"
	roomhandlers "goReact/webapp/admin/handlersAdmin/roomHandlers"
	seathandlers "goReact/webapp/admin/handlersAdmin/seatHandlers"
	usershandlers "goReact/webapp/admin/handlersAdmin/usersHandlers"
	"goReact/webapp/admin/middlewear/download"
	"net/http"
)

func (s *Server) configureRoutesAdmin() {
	//Auth
	s.router.Handle("GET", "/admin/login", auth.LoginAdmin(store.New(s.config)))
	s.router.Handle("POST", "/admin/auth", auth.AuthAdmin(store.New(s.config)))
	s.router.Handle("GET", "/admin/logout", auth.LogoutAdmin(store.New(s.config)))

	s.router.Handle("GET", "/admin/home", auth.HomeAdmin(store.New(s.config)))

	//User
	s.router.Handle("GET", "/admin/homeusers", usershandlers.HomeUsersHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/users", usershandlers.AllUsersHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/users/id/", usershandlers.GetUserByID(store.New(s.config)))
	s.router.Handle("GET", "/admin/users/csv/", usershandlers.PrintAllUsersCSV(store.New(s.config), download.DownloadFileHandler(store.New(s.config))))
	s.router.Handle("POST", "/admin/users/delete", usershandlers.DeleteUser(store.New(s.config)))
	s.router.Handle("POST", "/admin/users/update", usershandlers.UpdateUser(store.New(s.config)))
	s.router.Handle("POST", "/admin/users/new", usershandlers.NewUser(store.New(s.config)))

	//Hotel
	s.router.Handle("GET", "/admin/homehotels", hotelhandlers.HomeHotelHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/hotels", hotelhandlers.AllHotelsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/hotels/id", hotelhandlers.GetHotelByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/hotels/delete", hotelhandlers.DeleteHotels(store.New(s.config)))
	s.router.Handle("POST", "/admin/hotels/new", hotelhandlers.NewHotel(store.New(s.config)))
	s.router.Handle("POST", "/admin/hotels/update", hotelhandlers.UpdateHotel(store.New(s.config)))

	//Pet
	s.router.Handle("GET", "/admin/homepets", pethandlers.HomePetsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/pets", pethandlers.AllPetsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/pets/id", pethandlers.GetPetByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/pets/delete", pethandlers.DeletePets(store.New(s.config)))
	s.router.Handle("POST", "/admin/pets/new", pethandlers.NewPet(store.New(s.config)))
	s.router.Handle("POST", "/admin/pets/update", pethandlers.UpdatePet(store.New(s.config)))

	//Room
	s.router.Handle("GET", "/admin/homerooms", roomhandlers.HomeRoomHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/rooms", roomhandlers.AllRoomsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/rooms/id", roomhandlers.GetRoomByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/rooms/delete", roomhandlers.DeleteRooms(store.New(s.config)))
	s.router.Handle("POST", "/admin/rooms/new", roomhandlers.NewRoom(store.New(s.config)))
	s.router.Handle("POST", "/admin/rooms/update", roomhandlers.UpdateRoom(store.New(s.config)))

	//Seat
	s.router.Handle("GET", "/admin/homeseats", seathandlers.HomeSeatsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/seats", seathandlers.AllSeatsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/seats/id", seathandlers.GetSeatByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/seats/delete", seathandlers.DeleteSeats(store.New(s.config)))
	s.router.Handle("POST", "/admin/seats/new", seathandlers.NewSeat(store.New(s.config)))
	s.router.Handle("POST", "/admin/seats/update", seathandlers.UpdateSeat(store.New(s.config)))

	//Booking
	s.router.Handle("GET", "/admin/homebookings", bookinghandlers.HomeBookingHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/bookings", bookinghandlers.AllBookingsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/bookings/id", bookinghandlers.GetBookingByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/bookings/delete", bookinghandlers.DeleteBooking(store.New(s.config)))
	s.router.Handle("POST", "/admin/bookings/new", bookinghandlers.NewBooking(store.New(s.config)))
	s.router.Handle("POST", "/admin/bookings/update", bookinghandlers.UpdateBooking(store.New(s.config)))

	//Employee
	s.router.Handle("GET", "/admin/homeemployees", employeehandlers.HomeEmployeesHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/employees", employeehandlers.AllEmployeeHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/employees/id", employeehandlers.GetEmployeeByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/employees/delete", employeehandlers.DeleteEmployee(store.New(s.config)))
	s.router.Handle("POST", "/admin/employees/new", employeehandlers.NewEmployee(store.New(s.config)))
	s.router.Handle("POST", "/admin/employees/update", employeehandlers.UpdateEmployee(store.New(s.config)))

	//Permission
	s.router.Handle("GET", "/admin/permissions", permission.AllPermissons(store.New(s.config)))
	s.router.Handle("GET", "/admin/permissionsemployee", permission.GetPerByEmplID(store.New(s.config)))
	s.router.Handle("GET", "/admin/homepermissions", permission.HomePermissions(store.New(s.config)))
	s.router.Handle("GET", "/admin/addpermissions", permission.ShowAllPermissions(store.New(s.config)))
	s.router.Handle("GET", "/admin/permissionsemployyes", permission.AllPermissionsEmployees(store.New(s.config)))
	s.router.Handle("POST", "/admin/set", permission.AddPermissionsEmployee(store.New(s.config)))

	//Image
	s.router.Handle("GET", "/admin/homeimages", imagehandlers.HomeImageHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/images", imagehandlers.GetAllImagesHandle(store.New(s.config)))
	s.router.Handle("GET", "/admin/image/id", imagehandlers.GetImageHandle(store.New(s.config)))
	s.router.Handle("GET", "/admin/image/delete", imagehandlers.DeleteImageHandle(store.New(s.config)))
	s.router.Handle("POST", "/admin/image/upload", imagehandlers.SaveJPEGHandle(store.New(s.config)))
	s.router.Handle("POST", "/admin/image/update", imagehandlers.UpdateImage(store.New(s.config)))

	s.router.ServeFiles("/templates/*filepath", http.Dir("templates"))
}


