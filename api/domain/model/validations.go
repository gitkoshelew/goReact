package model

import (
	"errors"
	"regexp"
	"strings"

	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	latin    = regexp.MustCompile(`\p{Latin}`)
	cyrillic = regexp.MustCompile(`[\p{Cyrillic}]`)
	phone    = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
)

// WithoutSpaces ...
func WithoutSpaces(value interface{}) error {
	if strings.ContainsAny(value.(string), " ") {
		return errors.New("Field cannot contains spaces")
	}

	return nil
}

// IsLetterHyphenSpaces checks if string contains only letter(from simillar alphabet(latin or cyrillic)), hyphen or spaces
// Valid:"Name", "Name name", "Name-name"
// Invalid: "Name123", "NameИмя", "Name@name"
func IsLetterHyphenSpaces(value interface{}) error {
	s := value.(string)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "-", "", -1)

	err := is.UTFLetter.Validate(s)
	if err != nil {
		return errors.New("Only latin or cyrillic symblos, space and '-' symbol allowed")
	}
	if cyrillic.MatchString(s) && !latin.MatchString(s) {
		return nil
	} else if latin.MatchString(s) && !cyrillic.MatchString(s) {
		return nil
	}
	return errors.New("Only latin or cyrillic symblos, space and '-' symbol allowed")
}

// IsPhone ...
func IsPhone(value interface{}) error {
	s := value.(string)

	if phone.MatchString(s) {
		return nil
	}
	return errors.New("Invalid phone number format")
}

// IsRole checks if string matchs to a Role types of User
// ClientRole     = "client"
// EmployeeRole   = "employee"
// AnonymousRole  = "anonymous"
func IsRole(value interface{}) error {
	s := value.(Role)
	if s == ClientRole || s == EmployeeRole || s == AnonymousRole {
		return nil
	}
	return errors.New("Allowed roles for user: 'client', 'employee', 'anonymous'")
}

// IsSex checks if string matchs to a sex types of User
// Male    = "male"
// Female  = "female"
func IsSex(value interface{}) error {
	s := value.(Sex)
	if s == Male || s == Female {
		return nil
	}
	return errors.New("Allowed genders: 'male', 'female'")
}


// IsPetType checks if string matchs to a Pet types of Pets
// PetTypeCat = "cat"
// PetTypeDog = "dog"
func IsPetType(value interface{}) error {
	s := value.(PetType)
	if s == PetTypeCat || s == PetTypeDog {
		return nil
	}
	return errors.New("Allowed pet types: 'PetTypeCat', 'PetTypeDog'")
}

// IsPetType checks if string matchs to a Pet types of Pets
// PetTypeCat = "cat"
// PetTypeDog = "dog"
func IsEmployeePosition(value interface{}) error {
	s := value.(Position)
	if s == ManagerPosition || s == EmployeePosition || s == OwnerPosition || s == AdminPosition {
		return nil
	}
	return errors.New("fllowed pet types: 'ManagerPosition', 'EmployeePosition' ,'OwnerPosition','AdminPosition'")
}