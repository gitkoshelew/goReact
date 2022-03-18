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
