package model_test

import (
	"goReact/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.Login
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.Login {
				return model.TestLogin()
			},
			isValid: true,
		}, {
			name: "empty email",
			u: func() *model.Login {
				login := model.TestLogin()
				login.Email = ""
				return login
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.Login {
				login := model.TestLogin()
				login.Email = "invalid"
				return login
			},
			isValid: false,
		},
		{
			name: "SQL email",
			u: func() *model.Login {
				login := model.TestLogin()
				login.Email = "Sel--%^3 /** ecT"
				return login
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.Login {
				login := model.TestLogin()
				login.Password = ""
				return login
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.Login {
				login := model.TestLogin()
				login.Password = "1234"
				return login
			},
			isValid: false,
		},
		{
			name: "long password",
			u: func() *model.Login {
				login := model.TestLogin()
				login.Password = "1234567891012345678910123456789101234567891012345678910123456789101234567891012345678910123456789101234567891012345678910"
				return login
			},
			isValid: false,
		},
		{
			name: "sql password",
			u: func() *model.Login {
				login := model.TestLogin()
				login.Password = "ALt  9*/123#@! eR"
				return login
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
