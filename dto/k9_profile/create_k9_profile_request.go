package k9_profile

type CreateK9ProfileRequest struct {

	Name string `json:"name" binding:"required"`

	Breed string `json:"breed" binding:"required"`

	Gender string `json:"gender" binding:"required"`

	DateOfBirth string `json:"date_of_birth" binding:"required"`

	Microchip string `json:"microchip" binding:"required"`

	Color string `json:"color" binding:"required"`

	ServiceSince string `json:"service_since" binding:"required"` // Optional field, can be null
	
	ServiceEnd string `json:"service_end"` // Optional field, can be null
	
	Status int `json:"status" binding:"required"`
}