package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"image"
	"image/jpeg"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImagelRepository_SaveImage(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() (*model.ImageDTO, *image.Image)
		isValid bool
	}{
		{
			name: "valid",
			model: func() (*model.ImageDTO, *image.Image) {
				testStore.Open()
				imageDTO := model.TestImageDTO()
				file, _ := os.Open("./test_image.jpg")
				imageFile, _ := jpeg.Decode(file)
				return imageDTO, &imageFile
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, _, err := testStore.Image().SaveImage(tc.model())

				testStore.Image().Delete(*result)
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, _, err := testStore.Image().SaveImage(tc.model())
				testStore.Image().Delete(*result)
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestImagelRepository_Resize(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *image.Image
		isValid bool
	}{
		{
			name: "valid",
			model: func() *image.Image {
				testStore.Open()
				file, _ := os.Open("./test_image.jpg")
				imageFile, _ := jpeg.Decode(file)
				return &imageFile
			},
			isValid: true,
		},
		{
			name: "nil image",
			model: func() *image.Image {
				testStore.Open()
				return nil
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Image().ResizeImage(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Image().ResizeImage(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestImagelRepository_Delete(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		id      func() *int
		isValid bool
	}{
		{
			name: "valid",
			id: func() *int {
				testStore.Open()
				imageDTO := model.TestImageDTO()
				file, _ := os.Open("./test_image.jpg")
				imageFile, _ := jpeg.Decode(file)
				id, _, _ := testStore.Image().SaveImage(imageDTO, &imageFile)
				imageDTO.ImageID = *id
				return &imageDTO.ImageID
			},
			isValid: true,
		},
		{
			name: "invalid id",
			id: func() *int {
				testStore.Open()
				id := 0
				return &id
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Image().Delete(*tc.id())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.Image().Delete(*tc.id())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestImagelRepository_GetImageFromLocalStore(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.ImageDTO
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.ImageDTO {
				testStore.Open()
				imageDTO := model.TestImageDTO()
				file, _ := os.Open("./test_image.jpg")
				imageFile, _ := jpeg.Decode(file)
				id, _, _ := testStore.Image().SaveImage(imageDTO, &imageFile)
				imageDTO.ImageID = *id
				return imageDTO
			},
			isValid: true,
		},

		{
			name: "invalid id",
			model: func() *model.ImageDTO {
				testStore.Open()
				imageDTO := model.TestImageDTO()
				id := 0
				imageDTO.OwnerID = id
				return imageDTO
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Image().GetImageFromLocalStore(tc.model())
				testStore.Image().Delete(tc.model().ImageID)
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Image().GetImageFromLocalStore(tc.model())
				testStore.Image().Delete(tc.model().ImageID)
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}
