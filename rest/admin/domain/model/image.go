package model

import validation "github.com/go-ozzo/ozzo-validation"

// Image struct
type Image struct {
	ImageID int         `json:"imageId"`
	Type    ImageType   `json:"type"`
	OwnerID int         `json:"ownerId"`
	Format  ImageFormat `json:"format"`
}

// ImageDTO struct
type ImageDTO struct {
	ImageID int    `json:"imageId"`
	Type    string `json:"type"`
	OwnerID int    `json:"ownerId"`
	Format  string `json:"format"`
}

// ImageType ...
type ImageType string

// Image Types
var (
	PetImage     ImageType = "pet"
	RoomImage    ImageType = "room"
	UserImage    ImageType = "user"
	TestingImage ImageType = "test"
)

// ImageFormat ...
type ImageFormat string

// Image formats
var (
	FormatOriginal ImageFormat = "original"
	FormatQVGA     ImageFormat = "QVGA"
	FormatVGA      ImageFormat = "VGA"
	FormatHD720p   ImageFormat = "HD720p"
)

// Validate ...
func (i *ImageDTO) Validate() error {
	return validation.ValidateStruct(
		i,
		validation.Field(&i.Type, validation.Required, validation.By(IsImageType)),
		validation.Field(&i.OwnerID, validation.Required, validation.By(IsValidID)),
		validation.Field(&i.Format, validation.Required, validation.By(IsImageFormat)),
	)
}
