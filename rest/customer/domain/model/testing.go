package model

import "time"

// TestUser ...
func TestUser() *User {
	verified := true
	dateOfBirth := time.Time{}.AddDate(2000, 2, 2)
	return &User{
		UserID:      1,
		Email:       "email@example.org",
		Password:    "password",
		Role:        EmployeeRole,
		Verified:    &verified,
		Name:        "Name",
		Surname:     "Surname",
		MiddleName:  "MiddleName",
		Sex:         SexMale,
		DateOfBirth: &dateOfBirth,
		Address:     "Minsk Pr. Nezavisimosti 22-222",
		Phone:       "+375-29-154-89-33",
		Photo:       "Photo",
	}

}

// TestPet ...
func TestPet() *Pet {
	return &Pet{
		PetID:    1,
		Name:     "Name",
		Type:     PetTypeCat,
		Weight:   1,
		Diseases: "Disease",
		Owner:    *TestUser(),
		PhotoURL: "/",
	}
}
