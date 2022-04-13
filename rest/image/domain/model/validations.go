package model

import (
	"errors"
	"regexp"
	"strings"

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
	// ErrInvalidImageType ...
	ErrInvalidImageType = errors.New("invalid image type. Allowed image types: 'user', 'pet', 'room'")
	// ErrInvalidImageFormat ...
	ErrInvalidImageFormat = errors.New("invalid image format. Allowed image types: 'original', '', 'QVGA', 'VGA', 'HD720p'")
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

// IsImageType checks if string matchs to a Image types of images
//  PetImage   = "pet"
// 	RoomImage  = "room"
// 	UserImage  = "user"
func IsImageType(value interface{}) error {
	s := ImageType(value.(string))
	if s == PetImage || s == RoomImage || s == UserImage || s == TestingImage {
		return nil
	}
	return ErrInvalidImageType
}

// IsImageFormat checks if string matchs to a formats of images
// FormatOriginal  = "original"
// FormatQVGA      = "QVGA"
// FormatVGA       = "VGA"
// FormatHD720p    = "HD720p"
// empty string allowed - convert to original
func IsImageFormat(value interface{}) error {
	s := ImageFormat(value.(string))
	if s == FormatOriginal || s == FormatHD720p || s == FormatQVGA || s == FormatVGA || s == "" {
		return nil
	}
	return ErrInvalidImageFormat
}

// IsValidID ...
func IsValidID(value interface{}) error {
	id := value.(int)
	if id < 1 {
		return ErrInvalidID
	}
	return nil
}
