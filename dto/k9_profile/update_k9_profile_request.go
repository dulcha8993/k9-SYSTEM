package k9_profile

type UpdateK9ProfileRequest struct {

	Name *string `json:"name"`

	Breed *string `json:"breed"`	

	Gender *string `json:"gender"`

	DateOfBirth *string `json:"date_of_birth"`

	Microchip *string `json:"microchip"`

	Color *string `json:"color"`

	ServiceSince *string `json:"service_since"` // Optional field, can be null
	
	ServiceEnd *string `json:"service_end"` // Optional field, can be null
	
	Status *int `json:"status"`
}