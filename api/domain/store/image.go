package store

// ImageType ...
type ImageType string

// Image Types
var (
	PetImage  ImageType = "pet"
	RoomImage ImageType = "room"
)

// Image struct
type Image struct {
	ImageID int       `json:"imageId"`
	Type    ImageType `json:"type"`
	URL     string    `json:"url"`
	OwnerID int       `json:"ownerId"`
}
