package utils

import (
	"context"
	"fmt"
	"gateway/internal/apperror"
	"gateway/internal/client"
	"gateway/internal/client/customer"
	"gateway/internal/client/hotel"
	"io"
	"net/http"

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

var (
	// ErrNilPetID using AppError struct from apperror package
	ErrNilPetID = apperror.AppError{
		Message:          "Nil pet id is not allowed",
		DeveloperMessage: fmt.Sprintf("Nil pet id is not allowed. Cant use nil or 0 values while creating a new object"),
		Code:             fmt.Sprintf("%d", http.StatusBadRequest),
	}

	// ErrNilSeatID ...
	ErrNilSeatID = apperror.AppError{
		Message:          "Nil seat id is not allowed",
		DeveloperMessage: fmt.Sprintf("Nil seat id is not allowed. Cant use nil or 0 values while creating a new object"),
		Code:             fmt.Sprintf("%d", http.StatusBadRequest),
	}

	// ErrNilEmployeeID ...
	ErrNilEmployeeID = apperror.AppError{
		Message:          "Nil employee id is not allowed",
		DeveloperMessage: fmt.Sprintf("Nil employee id is not allowed. Cant use nil or 0 values while creating a new object"),
		Code:             fmt.Sprintf("%d", http.StatusBadRequest),
	}
)

// IsValid checks if data exist, type in keys object that you need to check and it's ID in value
// for keys use DataValidation constans
func IsValid(body io.Reader, nilAllowed bool, requiresValidation map[DataValidation]int, logger *logrus.Logger) error {

	for data := range requiresValidation {
		if requiresValidation[data] == 0 {
			delete(requiresValidation, data)
		}
	}

	id, ok := requiresValidation[PetValidation]
	if ok {
		logger.Debugf("Checking pet with id = %d", id)
		_, err := customer.Get(context.WithValue(context.Background(), client.CustomerGetQuerryParamsCtxKey, id), client.CustomerPetService, body)
		if err != nil {
			logger.Errorf("pet is not valid. error: %w", err)
			return err
		}
	} else if !ok && !nilAllowed {
		logger.Error(ErrNilPetID)
		return apperror.APIError(ErrNilPetID.Code, ErrNilPetID.Message, ErrNilPetID.DeveloperMessage)
	}

	id, ok = requiresValidation[EmployeeValidation]
	if ok {
		logger.Debugf("Checking employee with id = %d", id)
		_, err := hotel.Get(context.WithValue(context.Background(), client.HotelGetQuerryParamsCtxKey, id), client.HotelEmployeeService, body)
		if err != nil {
			logger.Errorf("employee is not valid. error: %w", err)
			return err
		}
	} else if !ok && !nilAllowed {
		logger.Error(ErrNilEmployeeID)
		return apperror.APIError(ErrNilEmployeeID.Code, ErrNilEmployeeID.Message, ErrNilEmployeeID.DeveloperMessage)
	}

	id, ok = requiresValidation[SeatValidation]
	if ok {
		logger.Debugf("Checking seat with id = %d", id)
		_, err := hotel.Get(context.WithValue(context.Background(), client.HotelGetQuerryParamsCtxKey, id), client.HotelSeatService, body)
		if err != nil {
			logger.Errorf("seat is not valid. error: %w", err)
			return err
		} else if !ok && !nilAllowed {
			logger.Error(ErrNilSeatID)
			return apperror.APIError(ErrNilSeatID.Code, ErrNilSeatID.Message, ErrNilSeatID.DeveloperMessage)
		}
	}
	return nil

}
