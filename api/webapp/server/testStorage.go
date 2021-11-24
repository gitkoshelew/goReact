package server

import (
	"goReact/domain/entity"
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
		entity.User{accounts[0], 1, "Ivan", "Ivanov", "Ivanovich", "22-1-1980", "Lenina 10", "80252552525", "ivan@mail.com"},
		entity.User{accounts[1], 2, "Sergey", "Sergeyev", "Sergeyvich", "13-5-1977", "Krasnaya 177", "8025123131", "sergey@mail.com"},
		entity.User{accounts[2], 3, "Vladimir", "Vladimirov", "Vladimirovich", "02-12-2000", "Vulica 30", "802521231235", "vladimir@mail.com"},
		entity.User{accounts[3], 4, "Olga", "Oleinikova", "Olegovna", "11-5-1991", "Nemiga 65", "8025944655", "olga@mail.com"},
		entity.User{accounts[4], 5, "Valeria", "Valerianovna", "Valerievna", "07-06-1987", "Bobruyskaya 7", "80252545645", "valeria@mail.com"},
	}
	return users
}
