package customer

import "time"

// UserDTO ...
type UserDTO struct {
	UserID      int       `json:"userId,omitempty"`
	Email       string    `json:"email,omitempty"`
	Role        string    `json:"role,omitempty"`
	Verified    bool      `json:"verified,omitempty"`
	Name        string    `json:"name,omitempty"`
	Surname     string    `json:"sName,omitempty"`
	MiddleName  string    `json:"mName,omitempty"`
	Sex         string    `json:"sex,omitempty"`
	DateOfBirth time.Time `json:"birthDate,omitempty"`
	Address     string    `json:"address,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Photo       string    `json:"photo,omitempty"`
}

// PetDTO ...
type PetDTO struct {
	PetID       int     `json:"petId,omitempty"`
	Name        string  `json:"name,omitempty"`
	Type        string  `json:"petType,omitempty"`
	Weight      float32 `json:"weight,omitempty"`
	Diseases    string  `json:"diseases,omitempty"`
	OwnerID     int     `json:"userId,omitempty"`
	PetPhotoURL string  `json:"petPhotoUrl,omitempty"`
}
