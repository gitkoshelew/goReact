package server

import (
	"goReact/domain/entity"
	"goReact/pkg/date"

	"cloud.google.com/go/civil"
)

// TestStorage ...
type TestStorage struct {
	Users entity.User
}

// GetAccounts ...
func GetAccounts() []entity.Account {

	accounts := []entity.Account{
		entity.Account{1, "login_1", "password_1"},
		entity.Account{2, "login_2", "password_2"},
		entity.Account{3, "login_3", "password_3"},
		entity.Account{4, "login_4", "password_4"},
		entity.Account{5, "login_5", "password_5"},
	}
	return accounts

}

// GetUsers ...
func GetUsers() []entity.User {
	accounts := GetAccounts()

	users := []entity.User{
		entity.User{accounts[0], 1, "Ivan", "Ivanov", "Ivanovich", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Lenina 10", "80252552525", "ivan@mail.com"},
		entity.User{accounts[1], 2, "Sergey", "Sergeyev", "Sergeyvich", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Krasnaya 177", "8025123131", "sergey@mail.com"},
		entity.User{accounts[2], 3, "Vladimir", "Vladimirov", "Vladimirovich", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Vulica 30", "802521231235", "vladimir@mail.com"},
		entity.User{accounts[3], 4, "Olga", "Oleinikova", "Olegovna", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Nemiga 65", "8025944655", "olga@mail.com"},
		entity.User{accounts[4], 5, "Valeria", "Valerianovna", "Valerievna", date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}, "Bobruyskaya 7", "80252545645", "valeria@mail.com"},
	}
	return users
}
