package k9_trainer

type K9TrainerUpdateRequest struct {

	Name *string `json:"name"`

	BadgeNumber *string `json:"badge_number"`

	Rank *string `json:"rank"`

	Email *string `json:"email"`
	
	Unit *string `json:"unit"`

	JoinDate *string `json:"join_date"`
	
	ExperienceYears *int `json:"experience_years"`

	Status *int `json:"status" gorm:"default:1"`

	Certifications []K9TrainerCertificationRequest `json:"certifications"`
}