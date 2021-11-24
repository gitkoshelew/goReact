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

func (s *Server) configureRouter() {
	s.router.HandlerFunc("GET", "/", s.handleHomePage())
	s.router.HandlerFunc("GET", "/users", s.handleUsers())
	s.router.HandlerFunc("POST", "/search_user", s.handleSearchUser())
}

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

func (s *Server) handleSearchUser() http.HandlerFunc {

	users := GetUsers()

	return func(w http.ResponseWriter, r *http.Request) {

		ID, err := strconv.Atoi(r.FormValue("ID"))
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			if user.UserID == ID {
				tmpl, err := template.ParseFiles("webapp/templates/users.html", "webapp/templates/header.html", "webapp/templates/footer.html")
				if err != nil {
					fmt.Fprintf(w, err.Error())
					log.Fatal(err)
				}
				tmpl.ExecuteTemplate(w, "users", [1]entity.User{user})

				http.Redirect(w, r, "/users", 301)
				return
			}
		}
		tmpl, err := template.ParseFiles("webapp/templates/users.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "users", [1]string{"User not foud"})

		http.Redirect(w, r, "/users", 301)
	}
}
