package utils

import (
	"context"
	"gateway/internal/client"
	"gateway/internal/client/customer"
	"gateway/internal/client/hotel"
	"io"

	"github.com/sirupsen/logrus"
)

// DataValidation ...
type DataValidation string

// DataValidation constants
const (
	PetValidation      DataValidation = "pet"
	EmployeeValidation DataValidation = "employee"
	SeatValidation     DataValidation = "seat"
)

// IsValid checks if data exist, type in keys object that you need to check and it's ID in value
// for keys use DataValidation constans
func IsValid(body io.Reader, requiresValidation map[DataValidation]int, logger *logrus.Logger) error {
	id, ok := requiresValidation[PetValidation]
	if ok {
		logger.Debugf("Checking pet with id = %d", id)
		_, err := customer.Get(context.WithValue(context.Background(), client.CustomerGetQuerryParamsCtxKey, id), client.CustomerPetService, body)
		if err != nil {
			logger.Errorf("pet is not valid. error: %w", err)
			return err
		}
	}

	id, ok = requiresValidation[EmployeeValidation]
	if ok {
		logger.Debugf("Checking employee with id = %d", id)
		_, err := hotel.Get(context.WithValue(context.Background(), client.HotelGetQuerryParamsCtxKey, id), client.HotelEmployeeService, body)
		if err != nil {
			logger.Errorf("employee is not valid. error: %w", err)
			return err
		}
	}

	id, ok = requiresValidation[SeatValidation]
	if ok {
		logger.Debugf("Checking seat with id = %d", id)
		_, err := hotel.Get(context.WithValue(context.Background(), client.HotelGetQuerryParamsCtxKey, id), client.HotelSeatService, body)
		if err != nil {
			logger.Errorf("seat is not valid. error: %w", err)
			return err
		}
	}
	return nil
}
