package model

// Image struct
type Image struct {
	ImageID int         `json:"imageId"`
	Type    ImageType   `json:"type"`
	URL     string      `json:"url"`
	OwnerID int         `json:"ownerId"`
	Format  ImageFormat `json:"format"`
}

// ImageDTO struct
type ImageDTO struct {
	ImageID int    `json:"imageId"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	OwnerID int    `json:"ownerId"`
	Format  string `json:"format"`
}

// ImageType ...
type ImageType string

// Image Types
var (
	PetImage  ImageType = "pet"
	RoomImage ImageType = "room"
	UserImage ImageType = "user"
)

// ImageFormat ...
type ImageFormat string

// Image Types
var (
	FormatOriginal ImageFormat = "original"
	FormatQVGA     ImageFormat = "QVGA"
	FormatVGA      ImageFormat = "VGA"
	FormatHD720p   ImageFormat = "HD720p"
)
