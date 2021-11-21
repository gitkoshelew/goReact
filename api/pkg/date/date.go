package date

import (
	"errors"
	"time"

	"cloud.google.com/go/civil"
)

// Date structure can be created like "date.Date{Date: civil.Date{Year: 2021, Month: 5, Day: 15}}"
// TODO: make creation of the struct one-level and don't expose civil library
type Date struct {
	civil.Date
}

func (d Date) Format(layout string) (string, error) {
	if !d.IsValid() {
		return "", errors.New("invalid date")
	}

	t, err := time.Parse("2006-01-02", d.String())
	if err != nil {
		return "", err
	}

	return t.Format(layout), nil
}
