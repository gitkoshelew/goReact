package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func ValidateId(id int, w http.ResponseWriter) error {
	if id < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct id")
		return errors.New("enter correct id")
	}
	return nil
}

func ValidateName(str string, w http.ResponseWriter) error {
	if len(str) < 2 && len(str) > 20 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct names")
		return errors.New("enter correct names")
	} else {
		for _, r := range str {
			if !unicode.IsLetter(r) {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Enter correct names")
				return errors.New("enter correct names")
			}
		}
	}
	return nil
}

func ValidateDateOfBirth(d time.Time, w http.ResponseWriter) error {
	now := time.Now()
	fmt.Println("now", now)
	fmt.Println(now.Year())
	s := now.Year() - d.Year()
	if s < 1 || s > 99 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct date of birth")
		return errors.New("enter correct date of birth")
	}
	return nil
}

func ValidateEmail(e string, w http.ResponseWriter) error {
	var emailRegexPattern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 100 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct email")
		return errors.New("enter correct email")
	}
	if !emailRegexPattern.MatchString(e) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct email")
		return errors.New("enter correct email")
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct email")
		return errors.New("enter correct email")
	}
	return nil
}

func ValidateWeight(weight int, w http.ResponseWriter) error {
	if weight < 0 && weight > 50 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct weight")
		return errors.New("enter correct weight")
	}
	return nil
}

func ValidateString(str string, w http.ResponseWriter) error {
	if len(strings.TrimSpace(str)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct string")
		return errors.New("enter correct string")
	}
	return nil
}

func ValidateCheckIn(d time.Time, w http.ResponseWriter) error {
	now := time.Now()
	if !d.After(now) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct check in date")
		return errors.New("enter correct check in date")
	}
	return nil
}

func ValidateCheckOut(in time.Time, out time.Time, w http.ResponseWriter) error {
	if !out.After(in) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Enter correct check out date")
		return errors.New("enter correct check out date")
	}
	return nil
}
