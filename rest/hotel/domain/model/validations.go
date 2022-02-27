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

// IsPetType checks if string matchs to a Pet types of Pets
// PetTypeCat = "cat"
// PetTypeDog = "dog"
func IsPetType(value interface{}) error {
	s := value.(PetType)
	if s == PetTypeCat || s == PetTypeDog {
		return nil
	}
	return errors.New("allowed pet types: 'PetTypeCat', 'PetTypeDog'")
}

// IsEmployeePosition checks if string matchs to a Employee Position
//	ManagerPosition  Position = "manager"
//	EmployeePosition Position = "employee"
//	OwnerPosition    Position = "owner"
//	AdminPosition    Position = "admin"
func IsEmployeePosition(value interface{}) error {
	s := value.(Position)
	if s == ManagerPosition || s == EmployeePosition || s == OwnerPosition || s == AdminPosition {
		return nil
	}
	return errors.New("allowed Employee Position: 'ManagerPosition', 'EmployeePosition' ,'OwnerPosition','AdminPosition'")
}
