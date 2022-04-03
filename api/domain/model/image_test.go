package model_test

import (
	"goReact/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImage_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		model   func() *model.ImageDTO
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.ImageDTO {
				return model.TestImageDTO()
			},
			isValid: true,
		},
		{
			name: "invalid image type",
			model: func() *model.ImageDTO {
				image := model.TestImageDTO()
				image.Type = "invalid type"
				return image
			},
			isValid: false,
		},
		{
			name: "invalid image format",
			model: func() *model.ImageDTO {
				image := model.TestImageDTO()
				image.Format = "invalid format"
				return image
			},
			isValid: false,
		},

		{
			name: "invalid Owner ID",
			model: func() *model.ImageDTO {
				image := model.TestImageDTO()
				image.OwnerID = 0
				return image
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.model().Validate())
			} else {
				assert.Error(t, tc.model().Validate())
			}
		})
	}
}
