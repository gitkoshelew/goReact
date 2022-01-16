package model_test

import (
	"goReact/domain/model"
	"testing"

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
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "1234"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = "Name@123"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Surname",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Surname = "Surname-Фамилия"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Surname",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Surname = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid MiddleName",
			u: func() *model.User {
				u := model.TestUser(t)
				u.MiddleName = "MiddleName %?№"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty MiddleName",
			u: func() *model.User {
				u := model.TestUser(t)
				u.MiddleName = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Role",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Role = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Sex",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Sex = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Phone",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Phone = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Phone",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Phone = ""
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

func TestUser_NewUser(t *testing.T) {
	u := model.TestUser(t)
	u.UserID = 1
	err := u.NewUser()

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUser_EncryptPassword(t *testing.T) {

	pass, err := model.EncryptPassword("password")
	assert.NoError(t, err)
	assert.NotNil(t, pass)
}

func TestUser_CheckPasswordHash(t *testing.T) {
	password := "password"
	encrypt, _ := model.EncryptPassword(password)

	err := model.CheckPasswordHash(encrypt, password)
	assert.NoError(t, err)

	err = model.CheckPasswordHash(encrypt, "AnotherPassword")
	assert.Error(t, err)

	err = model.CheckPasswordHash(encrypt, "")
	assert.Error(t, err)

	err = model.CheckPasswordHash("encrypt", password)
	assert.Error(t, err)

	err = model.CheckPasswordHash("", password)
	assert.Error(t, err)

	err = model.CheckPasswordHash("", "")
	assert.Error(t, err)
}
