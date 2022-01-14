package model

import (
	"testing"
	"time"
)

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:       "email@example.org",
		Password:    "password",
		Role:        ClientRole,
		Verified:    true,
		Name:        "Name",
		Surname:     "Surname",
		MiddleName:  "MiddleName",
		Sex:         Male,
		DateOfBirth: time.Time{}.AddDate(2000, 2, 2),
		Address:     "Minsk Pr. Nezavisimosti 22-222",
		Phone:       "+375-29-154-89-33",
		Photo:       "Photo",
	}

}

// Testing instance of hotel
func TestHotel(t *testing.T) *Hotel {
	return &Hotel{
		HotelID: 1,
		Name:    "Name",
		Address: "Minsk ul sovetskaya 18",
	}
}
