package server

import (
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// Server ...
type Server struct {
	config *webapp.Config
	logger *log.Logger
	router *httprouter.Router
}

// New ...
func New(config *webapp.Config) *Server {
	return &Server{
		config: config,
		logger: log.New(os.Stdout, "http: ", log.Ldate|log.Ltime|log.Lmsgprefix),
		router: httprouter.New(),
	}
}

// Start ...
func (s *Server) Start() error {

	s.configureRouter()

	s.logger.Printf("Server starting ...")

	return http.ListenAndServe(s.config.ServerAddress(), s.router)
}

// Server ...
func (s *Server) configureRouter() {
	s.router.HandlerFunc("GET", "/", s.handleHomePage())

	s.router.HandlerFunc("GET", "/accounts", s.handleAccounts())
	s.router.HandlerFunc("GET", "/account", s.handleAccountSearch())

	s.router.HandlerFunc("GET", "/users", s.handleUsers())
	s.router.HandlerFunc("GET", "/user", s.handleUserSearch())

	s.router.HandlerFunc("GET", "/employees", s.handleEmployees())
	s.router.HandlerFunc("GET", "/employee", s.handleEmployeeSearch())

	s.router.HandlerFunc("GET", "/clients", s.handleClients())
	s.router.HandlerFunc("GET", "/client", s.handleClientSearch())

	s.router.HandlerFunc("GET", "/pets", s.handlePets())
	s.router.HandlerFunc("GET", "/pet", s.handlePetSearch())

	s.router.HandlerFunc("GET", "/hotels", s.handleHotels())
	s.router.HandlerFunc("GET", "/hotel", s.handleHotelSearch())

	s.router.HandlerFunc("GET", "/rooms", s.handleHotelRooms())
	s.router.HandlerFunc("GET", "/room", s.handleHotelRoomSearch())

	s.router.HandlerFunc("GET", "/seats", s.handleHotelRoomSeats())
	s.router.HandlerFunc("GET", "/seat", s.handleHotelRoomSeatSearch())

	s.router.HandlerFunc("GET", "/bookings", s.handleBookings())
	s.router.HandlerFunc("GET", "/booking", s.handleBookingSearch())
}

// handleHomePage opens a main page, URL: "/"
func (s *Server) handleHomePage() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/home_page.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "home_page", nil)
	}
}

// handleAccounts opens an account page, URL: "/accounts". Shows all accounts, can search one by id
func (s *Server) handleAccounts() http.HandlerFunc {

	accounts := GetAccounts()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/accounts.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "accounts", accounts)
	}
}

// handleAccountSearch shows an account by id, URL"/account?id="
func (s *Server) handleAccountSearch() http.HandlerFunc {

	accounts := GetAccounts()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var account entity.Account
		accountFound := false

		for _, a := range accounts {
			if a.AccountID == id {
				account = a
				accountFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_account.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if accountFound {
			tmpl.ExecuteTemplate(w, "show_account", account)
		} else {
			tmpl.ExecuteTemplate(w, "show_account", "Account not found")
		}

	}
}

// handleUsersopens an user page, URL: "/users". Shows all user, can search one by id
func (s *Server) handleUsers() http.HandlerFunc {

	users := GetUsers()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/users.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "users", users)
	}
}

// handleUserSearch shows an user by id, URL"/user?id="
func (s *Server) handleUserSearch() http.HandlerFunc {

	users := GetUsers()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var user entity.User
		userFound := false

		for _, a := range users {
			if a.UserID == id {
				user = a
				userFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_user.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if userFound {
			tmpl.ExecuteTemplate(w, "show_user", user)
		} else {
			tmpl.ExecuteTemplate(w, "show_user", "User not found")
		}

	}
}

// handleEmployees opens an employee page, URL: "/employees". Shows all employees, can search one by id
func (s *Server) handleEmployees() http.HandlerFunc {

	employees := GetEmployees()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/employees.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "employees", employees)
	}
}

// handleEmployeeSearch shows an employee by id, URL"/employee?id="
func (s *Server) handleEmployeeSearch() http.HandlerFunc {

	employees := GetEmployees()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var employee entity.Employee
		employeeFound := false

		for _, a := range employees {
			if a.EmployeeID == id {
				employee = a
				employeeFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_employee.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if employeeFound {
			tmpl.ExecuteTemplate(w, "show_employee", employee)
		} else {
			tmpl.ExecuteTemplate(w, "show_employee", "Employee not found")
		}

	}
}

// handleClients opens a client page, URL: "/clients". Shows all accounts, can search one by id
func (s *Server) handleClients() http.HandlerFunc {

	clients := GetClients()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/clients.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "clients", clients)
	}
}

// handleClientSearch shows a client by id, URL"/client?id="
func (s *Server) handleClientSearch() http.HandlerFunc {

	clients := GetClients()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var client entity.Client
		clientFound := false

		for _, a := range clients {
			if a.ClientID == id {
				client = a
				clientFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_client.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if clientFound {
			tmpl.ExecuteTemplate(w, "show_client", client)
		} else {
			tmpl.ExecuteTemplate(w, "show_client", "Client not found")
		}
	}
}

// handlePets opens a pet page, URL: "/pets". Shows all pets, can search one by id
func (s *Server) handlePets() http.HandlerFunc {

	pets := GetPets()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/pets.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "pets", pets)
	}
}

// handlePetSearch shows a pet by id, URL"/pet?id="
func (s *Server) handlePetSearch() http.HandlerFunc {

	pets := GetPets()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var pet entity.Pet
		petFound := false

		for _, a := range pets {
			if a.PetID == id {
				pet = a
				petFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_pet.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if petFound {
			tmpl.ExecuteTemplate(w, "show_pet", pet)
		} else {
			tmpl.ExecuteTemplate(w, "show_pet", "Pet not found")
		}

	}
}

// handleHotels opens a hotel page, URL: "/hotels". Shows all hotels, can search one by id
func (s *Server) handleHotels() http.HandlerFunc {

	hotels := GetHotels()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/hotels.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "hotels", hotels)
	}
}

// handleHotelSearch shows a hotel by id, URL"/hotel?id="
func (s *Server) handleHotelSearch() http.HandlerFunc {

	hotels := GetHotels()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var hotel entity.Hotel
		hotelFound := false

		for _, a := range hotels {
			if a.HotelID == id {
				hotel = a
				hotelFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_hotel.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if hotelFound {
			tmpl.ExecuteTemplate(w, "show_hotel", hotel)
		} else {
			tmpl.ExecuteTemplate(w, "show_hotel", "Hotel not found")
		}
	}
}

// handleHotelRooms opens an hotel rooms page, URL: "/rooms". Shows all hotel rooms, can search one by id
func (s *Server) handleHotelRooms() http.HandlerFunc {

	rooms := GetHotelRooms()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/rooms.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "rooms", rooms)
	}
}

// handleHotelRoomSearch shows a hotel room by id, URL"/room?id="
func (s *Server) handleHotelRoomSearch() http.HandlerFunc {

	rooms := GetHotelRooms()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var room entity.HotelRoom
		roomFound := false

		for _, a := range rooms {
			if a.HotelRoomID == id {
				room = a
				roomFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_room.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if roomFound {
			tmpl.ExecuteTemplate(w, "show_room", room)
		} else {
			tmpl.ExecuteTemplate(w, "show_room", "Room not found")
		}
	}
}

// handleHotelRoomSeats opens an hotel room seats page, URL: "/seats". Shows all hotel room seats, can search one by id
func (s *Server) handleHotelRoomSeats() http.HandlerFunc {

	hotelRoomSeats := GetHotelRoomSeats()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/seats.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "seats", hotelRoomSeats)
	}
}

// handleHotelRoomSeatSearch shows an hotel room seats by id, URL"/seat?id="
func (s *Server) handleHotelRoomSeatSearch() http.HandlerFunc {

	seats := GetHotelRoomSeats()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var seat entity.HotelRoomSeat
		seatFound := false

		for _, a := range seats {
			if a.HotelRoomSeatID == id {
				seat = a
				seatFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_seat.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if seatFound {
			tmpl.ExecuteTemplate(w, "show_seat", seat)
		} else {
			tmpl.ExecuteTemplate(w, "show_seat", "Seat not found")
		}
	}
}

// handleBookings opens an booking page, URL: "/bookings". Shows all bookings, can search one by id
func (s *Server) handleBookings() http.HandlerFunc {

	bookings := GetBookings()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/bookings.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "bookings", bookings)
	}
}

// handleBookingSearch shows an booking by id, URL"/booking?id="
func (s *Server) handleBookingSearch() http.HandlerFunc {

	bookings := GetBookings()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var booking entity.Booking
		bookingFound := false

		for _, a := range bookings {
			if a.BookingID == id {
				booking = a
				bookingFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_booking.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if bookingFound {
			tmpl.ExecuteTemplate(w, "show_booking", booking)
		} else {
			tmpl.ExecuteTemplate(w, "show_booking", "Booking not found")
		}
	}
}
