package model_test

import (
	"goReact/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser()
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser()
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser()
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "SQL email",
			u: func() *model.User {
				u := model.TestUser()
				u.Email = "Sel--%^3 /** ecT"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = "1234"
				return u
			},
			isValid: false,
		},
		{
			name: "long password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = "1234567891012345678910123456789101234567891012345678910123456789101234567891012345678910123456789101234567891012345678910"
				return u
			},
			isValid: false,
		},
		{
			name: "sql password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = "ALt  9*/123#@! eR"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Role",
			u: func() *model.User {
				u := model.TestUser()
				u.Role = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Role",
			u: func() *model.User {
				u := model.TestUser()
				u.Role = ""
				return u
			},
			isValid: true,
		},
		{
			name: "SQL Role",
			u: func() *model.User {
				u := model.TestUser()
				u.Role = "ALt  --__- eR"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Name",
			u: func() *model.User {
				u := model.TestUser()
				u.Name = "Name@123"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Name",
			u: func() *model.User {
				u := model.TestUser()
				u.Name = ""
				return u
			},
			isValid: false,
		},
		{
			name: "SQL Name",
			u: func() *model.User {
				u := model.TestUser()
				u.Name = "AlT*&^er"
				return u
			},
			isValid: false,
		},
		{
			name: "Short Name",
			u: func() *model.User {
				u := model.TestUser()
				u.Name = "a"
				return u
			},
			isValid: false,
		},
		{
			name: "Long Name",
			u: func() *model.User {
				u := model.TestUser()
				u.Name = "adsadasdasdasdasdsadasdaSDadADSdasasdasdsadaddadas"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Surname",
			u: func() *model.User {
				u := model.TestUser()
				u.Surname = "Surname-Фамилия"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Surname",
			u: func() *model.User {
				u := model.TestUser()
				u.Surname = ""
				return u
			},
			isValid: false,
		},
		{
			name: "SQL Surname",
			u: func() *model.User {
				u := model.TestUser()
				u.Surname = "AlT*&^er"
				return u
			},
			isValid: false,
		},
		{
			name: "Short Surname",
			u: func() *model.User {
				u := model.TestUser()
				u.Surname = "a"
				return u
			},
			isValid: false,
		},
		{
			name: "Long Surname",
			u: func() *model.User {
				u := model.TestUser()
				u.Surname = "adsadasdasdasdasdsadasdaSDadADSdasasdasdsadaddadas"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid MiddleName",
			u: func() *model.User {
				u := model.TestUser()
				u.MiddleName = "MiddleName %?№"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty MiddleName",
			u: func() *model.User {
				u := model.TestUser()
				u.MiddleName = ""
				return u
			},
			isValid: false,
		},
		{
			name: "SQL MiddleName",
			u: func() *model.User {
				u := model.TestUser()
				u.MiddleName = "AlT*&^er"
				return u
			},
			isValid: false,
		},
		{
			name: "Long MiddleName",
			u: func() *model.User {
				u := model.TestUser()
				u.MiddleName = "adsadasdasdasdasdsadasdaSDadADSdasasdasdsadaddadas"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Sex",
			u: func() *model.User {
				u := model.TestUser()
				u.Sex = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Sex",
			u: func() *model.User {
				u := model.TestUser()
				u.Sex = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Under age 18 DateOfBirth",
			u: func() *model.User {
				u := model.TestUser()
				dateOfBirth := time.Now()
				u.DateOfBirth = &dateOfBirth
				return u
			},
			isValid: false,
		},
		{
			name: "Above age 100 DateOfBirth",
			u: func() *model.User {
				u := model.TestUser()
				dateOfBirth := time.Now().AddDate(-100, 0, -1)
				u.DateOfBirth = &dateOfBirth
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Address",
			u: func() *model.User {
				u := model.TestUser()
				u.Address = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Short Address",
			u: func() *model.User {
				u := model.TestUser()
				u.Address = "asd"
				return u
			},
			isValid: false,
		},
		{
			name: "Long Address",
			u: func() *model.User {
				u := model.TestUser()
				u.Address = "asdasdasdsadasdas sdad asdasdas dasdasd asdasd asdas dsa dasd"
				return u
			},
			isValid: false,
		},
		{
			name: "SQL Address",
			u: func() *model.User {
				u := model.TestUser()
				u.Address = "AL*6789 _-=--t=@#!#er"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Phone",
			u: func() *model.User {
				u := model.TestUser()
				u.Phone = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Phone",
			u: func() *model.User {
				u := model.TestUser()
				u.Phone = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Photo",
			u: func() *model.User {
				u := model.TestUser()
				u.Photo = ""
				return u
			},
			isValid: true,
		},
		{
			name: "SQL Photo",
			u: func() *model.User {
				u := model.TestUser()
				u.Photo = "AlTE@##4r"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_EncryptPassword(t *testing.T) {
	password := "password"

	pass, err := model.EncryptPassword(password)
	assert.NoError(t, err)
	assert.NotNil(t, pass)

	assert.NotEqual(t, pass, password)
}

func TestUser_CheckPasswordHash(t *testing.T) {
	password := "password"
	encryptedPassword, _ := model.EncryptPassword(password)

	t.Run("TestUser_CheckPasswordHash: Valid", func(t *testing.T) {
		err := model.CheckPasswordHash(encryptedPassword, password)
		assert.NoError(t, err)
	})

	t.Run("TestUser_CheckPasswordHash: Initial password is invalid", func(t *testing.T) {
		anotherPassword := "Another Password"
		err := model.CheckPasswordHash(encryptedPassword, anotherPassword)
		assert.Error(t, err)
	})

	t.Run("TestUser_CheckPasswordHash: Encrypted password is invalid", func(t *testing.T) {
		anotherEncryptedPassword := "another encryptedPassword"
		err := model.CheckPasswordHash(encryptedPassword, anotherEncryptedPassword)
		assert.Error(t, err)
	})
}
