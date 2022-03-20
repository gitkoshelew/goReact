package model_test

import (
	"goReact/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotel_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		h       func() *model.Hotel
		isValid bool
	}{
		{
			name: "valid",
			h: func() *model.Hotel {
				return model.TestHotel()
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.h().Validate())
			} else {
				assert.Error(t, tc.h().Validate())
			}
		})
	}
}
