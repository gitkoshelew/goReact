package model_test

import (
	"hotel/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmployee_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		e       func() *model.EmployeeDTO
		isValid bool
	}{
		{
			name: "valid",
			e: func() *model.EmployeeDTO {
				return model.TestEmployeeDTO()
			},
			isValid: true,
		},
		{
			name: "nil employee",
			e: func() *model.EmployeeDTO {
				return &model.EmployeeDTO{}
			},
			isValid: false,
		},
		{
			name: "invalid HotelID",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.HotelID = 0
				return employee
			},
			isValid: false,
		},
		{
			name: "invalid Position",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Position = "invalid position"
				return employee
			},
			isValid: false,
		},
		{
			name: "empty email",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Email = ""
				return employee
			},
			isValid: false,
		},
		{
			name: "invalid email",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Email = "invalid"
				return employee
			},
			isValid: false,
		},
		{
			name: "SQL email",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Email = "Sel--%^3 /** ecT"
				return employee
			},
			isValid: false,
		},
		{
			name: "Invalid Role",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Role = "Invalid"
				return employee
			},
			isValid: false,
		},
		{
			name: "Empty Role",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Role = ""
				return employee
			},
			isValid: true,
		},
		{
			name: "SQL Role",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Role = "ALt  --__- eR"
				return employee
			},
			isValid: false,
		},
		{
			name: "Invalid Name",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Name = "Name@123"
				return employee
			},
			isValid: false,
		},
		{
			name: "Empty Name",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Name = ""
				return employee
			},
			isValid: false,
		},
		{
			name: "SQL Name",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Name = "AlT*&^er"
				return employee
			},
			isValid: false,
		},
		{
			name: "Short Name",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Name = "a"
				return employee
			},
			isValid: false,
		},
		{
			name: "Long Name",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Name = "adsadasdasdasdasdsadasdaSDadADSdasasdasdsadaddadas"
				return employee
			},
			isValid: false,
		},
		{
			name: "Invalid Surname",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Surname = "Surname-Фамилия"
				return employee
			},
			isValid: false,
		},
		{
			name: "Empty Surname",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Surname = ""
				return employee
			},
			isValid: false,
		},
		{
			name: "SQL Surname",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Surname = "AlT*&^er"
				return employee
			},
			isValid: false,
		},
		{
			name: "Short Surname",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Surname = "a"
				return employee
			},
			isValid: false,
		},
		{
			name: "Long Surname",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Surname = "adsadasdasdasdasdsadasdaSDadADSdasasdasdsadaddadas"
				return employee
			},
			isValid: false,
		},
		{
			name: "Invalid MiddleName",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.MiddleName = "MiddleName %?№"
				return employee
			},
			isValid: false,
		},
		{
			name: "Empty MiddleName",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.MiddleName = ""
				return employee
			},
			isValid: false,
		},
		{
			name: "SQL MiddleName",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.MiddleName = "AlT*&^er"
				return employee
			},
			isValid: false,
		},
		{
			name: "Long MiddleName",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.MiddleName = "adsadasdasdasdasdsadasdaSDadADSdasasdasdsadaddadas"
				return employee
			},
			isValid: false,
		},
		{
			name: "Invalid Sex",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Sex = "Invalid"
				return employee
			},
			isValid: false,
		},
		{
			name: "Empty Sex",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Sex = "Invalid"
				return employee
			},
			isValid: false,
		},
		{
			name: "Under age 18 DateOfBirth",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				dateOfBirth := time.Now()
				employee.DateOfBirth = &dateOfBirth
				return employee
			},
			isValid: false,
		},
		{
			name: "Above age 100 DateOfBirth",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				dateOfBirth := time.Now().AddDate(-100, 0, -1)
				employee.DateOfBirth = &dateOfBirth
				return employee
			},
			isValid: false,
		},
		{
			name: "Empty Address",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Address = ""
				return employee
			},
			isValid: false,
		},
		{
			name: "Short Address",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Address = "asd"
				return employee
			},
			isValid: false,
		},
		{
			name: "Long Address",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Address = "asdasdasdsadasdas sdad asdasdas dasdasd asdasd asdas dsa dasd"
				return employee
			},
			isValid: false,
		},
		{
			name: "SQL Address",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Address = "AL*6789 _-=--t=@#!#er"
				return employee
			},
			isValid: false,
		},
		{
			name: "Invalid Phone",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Phone = "Invalid"
				return employee
			},
			isValid: false,
		},
		{
			name: "Empty Phone",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Phone = ""
				return employee
			},
			isValid: false,
		},
		{
			name: "Empty Photo",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Photo = ""
				return employee
			},
			isValid: true,
		},
		{
			name: "SQL Photo",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Photo = "AlTE@##4r"
				return employee
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.e().Validate())
			} else {
				assert.Error(t, tc.e().Validate())
			}
		})
	}
}
