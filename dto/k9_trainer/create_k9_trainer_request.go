package k9_trainer

type K9TrainerRequest struct {

	Name string `json:"name" binding:"required"`

	BadgeNumber string `json:"badge_number" binding:"required"`

	Rank string `json:"rank" binding:"required"`

	Email string `json:"email" binding:"required"`
	
	Unit string `json:"unit" binding:"required"`

	JoinDate string `json:"join_date"`
	
	ExperienceYears int `json:"experience_years"`

	Status int `json:"status" gorm:"default:1"`

	Certifications []K9TrainerCertificationRequest `json:"certifications"`
}	
