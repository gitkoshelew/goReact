package model

import (
	"errors"
	"regexp"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	// regex
	latin       = regexp.MustCompile(`\p{Latin}`)
	cyrillic    = regexp.MustCompile(`[\p{Cyrillic}]`)
	phone       = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	noSQL       = regexp.MustCompile(`\b(ALTER|CREATE|DELETE|DROP|EXEC(UTE){0,1}|INSERT( +INTO){0,1}|MERGE|SELECT|UPDATE|UNION( +ALL){0,1})\b`)
	onlyLetters = regexp.MustCompile("[^a-zA-Zа-яА-Я]+")

	// errors

	// ErrContainsSQL ...
	ErrContainsSQL = errors.New("no SQL commands allowed to input")
	// ErrInvalidBithDate ...
	ErrInvalidBithDate = errors.New("invalid date of birth. Age must be from 18 to 100. Date format RFC3339")
	// ErrInvalidBookingStatus ...
	ErrInvalidBookingStatus = errors.New("invalid booking status. Allowed Booking Statuses: 'pending', 'in-progress' ,'completed','cancelled'")
	// ErrInvalidEmployeePosition ...
	ErrInvalidEmployeePosition = errors.New("invalid employee position. Allowed pet types: 'manager', 'employee' ,'owner','admin'")
	// ErrInvalidPetType ...
	ErrInvalidPetType = errors.New("invalid pet type. Allowed pet types: 'cat', 'dog'")
	// ErrInvalidSexType ...
	ErrInvalidSexType = errors.New("invalid gender. Allowed genders: 'male', 'female'")
	// ErrInvalidUserRole ...
	ErrInvalidUserRole = errors.New("invalid roles. Allowed roles for user: 'client', 'employee', 'anonymous'")
	// ErrInvalidPhoneNumber ...
	ErrInvalidPhoneNumber = errors.New("invalid phone number format")
	// ErrInvalidAlphabet ...
	ErrInvalidAlphabet = errors.New("only latin or cyrillic symblos allowed")
	// ErrInvalidSymbol ...
	ErrInvalidSymbol = errors.New("invalid symbol used. Only space and '-' symbols allowed")
	// ErrInvalidStartDate ...
	ErrInvalidStartDate = errors.New("invalid start date. Start date cannot be before today")
	// ErrInvalidEndDate ...
	ErrInvalidEndDate = errors.New("invalid end date. End date cannot be before today")
	// ErrInvalidID ...
	ErrInvalidID = errors.New("invalid input: id")
)

// IsLetterHyphenSpaces checks if string contains only letter(from simillar alphabet(latin or cyrillic)), hyphen or spaces
// Valid:"Name", "Name name", "Name-name"
// Invalid: "Name123", "NameИмя", "Name@name"
func IsLetterHyphenSpaces(value interface{}) error {
	s := value.(string)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "-", "", -1)

	err := is.UTFLetter.Validate(s)
	if err != nil {
		return ErrInvalidSymbol
	}
	if cyrillic.MatchString(s) && !latin.MatchString(s) {
		return nil
	} else if latin.MatchString(s) && !cyrillic.MatchString(s) {
		return nil
	}
	return ErrInvalidAlphabet
}

// IsPhone ...
func IsPhone(value interface{}) error {
	s := value.(string)

	if phone.MatchString(s) {
		return nil
	}
	return ErrInvalidPhoneNumber
}

// IsRole checks if string matchs to a Role types of User
// ClientRole     = "client"
// EmployeeRole   = "employee"
// AnonymousRole  = "anonymous"
// if Role with be nil - wont return error
func IsRole(value interface{}) error {
	s := Role(value.(string))
	if s == ClientRole || s == EmployeeRole || s == "" {
		return nil
	}
	return ErrInvalidUserRole
}

// IsSex checks if string matchs to a sex types of User
// Male    = "male"
// Female  = "female"
func IsSex(value interface{}) error {
	s := Sex(value.(string))
	if s == SexMale || s == SexFemale {
		return nil
	}
	return ErrInvalidSexType
}

// IsPetType checks if string matchs to a Pet types of Pets
// PetTypeCat = "cat"
// PetTypeDog = "dog"
func IsPetType(value interface{}) error {
	s := PetType(value.(string))
	if s == PetTypeCat || s == PetTypeDog {
		return nil
	}
	return ErrInvalidPetType
}

// IsEmployeePosition ...
func IsEmployeePosition(value interface{}) error {
	s := Position(value.(string))
	if s == ManagerPosition || s == EmployeePosition || s == OwnerPosition || s == AdminPosition {
		return nil
	}
	return ErrInvalidEmployeePosition
}

// IsValidBirthDate ...
func IsValidBirthDate(value interface{}) error {
	t := time.Now()
	d := value.(*time.Time)
	err := validation.Validate(d.Format(time.RFC3339), validation.Date(time.RFC3339).Max(t.AddDate(-18, 0, 0)).Min(t.AddDate(-100, 0, 0)))
	if err != nil {
		return ErrInvalidBithDate
	}
	return nil
}

// IsSQL ...
func IsSQL(value interface{}) error {
	s := value.(string)

	if noSQL.MatchString(strings.ToUpper(s)) {
		return ErrContainsSQL
	}

	str := onlyLetters.ReplaceAllString(s, "")

	if noSQL.MatchString(strings.ToUpper(str)) {
		return ErrContainsSQL
	}
	return nil
}

// IsValidStartDate ...
func IsValidStartDate(value interface{}) error {
	t := time.Now()
	d := value.(*time.Time)
	err := validation.Validate(d.Format(time.RFC3339), validation.Date(time.RFC3339).Min(t.AddDate(0, 0, -1)))
	if err != nil {
		return ErrInvalidStartDate
	}
	return nil
}

// IsValidEndDate ...
func IsValidEndDate(value interface{}) error {
	t := time.Now()
	d := value.(*time.Time)
	err := validation.Validate(d.Format(time.RFC3339), validation.Date(time.RFC3339).Min(t))
	if err != nil {
		return ErrInvalidEndDate
	}
	return nil
}

// IsValidID ...
func IsValidID(value interface{}) error {
	id := value.(int)
	if id < 1 {
		return ErrInvalidID
	}
	return nil
}
