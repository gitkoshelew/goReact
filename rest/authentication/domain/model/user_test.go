package model_test

import (
	"auth/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.UserDTO
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.UserDTO {
				return model.TestUserDTO()
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "SQL email",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Email = "Sel--%^3 /** ecT"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Password = "1234"
				return u
			},
			isValid: false,
		},
		{
			name: "long password",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Password = "1234567891012345678910123456789101234567891012345678910123456789101234567891012345678910123456789101234567891012345678910"
				return u
			},
			isValid: false,
		},
		{
			name: "sql password",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Password = "ALt  9*/123#@! eR"
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Role",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Role = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Role",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Role = ""
				return u
			},
			isValid: false,
		},
		{
			name: "SQL Role",
			u: func() *model.UserDTO {
				u := model.TestUserDTO()
				u.Role = "ALt  --__- eR"
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
