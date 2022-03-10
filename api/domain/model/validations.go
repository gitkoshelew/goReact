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
		return errors.New("field cannot contains spaces")
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
		return errors.New("only latin or cyrillic symblos, space and '-' symbol allowed")
	}
	if cyrillic.MatchString(s) && !latin.MatchString(s) {
		return nil
	} else if latin.MatchString(s) && !cyrillic.MatchString(s) {
		return nil
	}
	return errors.New("only latin or cyrillic symblos, space and '-' symbol allowed")
}

// IsPhone ...
func IsPhone(value interface{}) error {
	s := value.(string)

	if phone.MatchString(s) {
		return nil
	}
	return errors.New("invalid phone number format")
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
	return errors.New("allowed roles for user: 'client', 'employee', 'anonymous'")
}

// IsSex checks if string matchs to a sex types of User
// Male    = "male"
// Female  = "female"
func IsSex(value interface{}) error {
	s := value.(Sex)
	if s == SexMale || s == SexFemale {
		return nil
	}
	return errors.New("allowed genders: 'male', 'female'")
}

// IsPetType checks if string matchs to a Pet types of Pets
// PetTypeCat = "cat"
// PetTypeDog = "dog"
func IsPetType(value interface{}) error {
	s := value.(PetType)
	if s == "cat" || s == "dog" {
		return nil
	}
	return errors.New("allowed pet types: 'PetTypeCat', 'PetTypeDog'")
}

// IsEmployeePosition ...
func IsEmployeePosition(value interface{}) error {
	s := value.(Position)
	if s == ManagerPosition || s == EmployeePosition || s == OwnerPosition || s == AdminPosition {
		return nil
	}
	return errors.New("allowed pet types: 'ManagerPosition', 'EmployeePosition' ,'OwnerPosition','AdminPosition'")
}

// IsBookingStatus checks if string matchs to a BookingStatus
// BookingStatusPending    BookingStatus = "pending"
// BookingStatusInProgress BookingStatus = "in-progress"
// BookingStatusCompleted  BookingStatus = "completed"
// BookingStatusCancelled  BookingStatus = "cancelled"
func IsBookingStatus(value interface{}) error {
	s := value.(BookingStatus)
	if s == "pending" || s == "in-progress" || s == "completed" || s == "cancelled" {
		return nil
	}
	return errors.New("allowed Booking Status: 'BookingStatusPending', 'BookingStatusInProgress' ,'BookingStatusCompleted','BookingStatusCancelled'")
}
