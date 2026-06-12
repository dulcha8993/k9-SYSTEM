package k9_trainer

import (
	"time"

	"github.com/google/uuid"
)

// import (
// 	TrainingRecord "k9-system/models/training_record"
// )

type K9Trainer struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Name string `gorm:"not null"`

	BadgeNumber string `gorm:"not null"`

	Rank string `gorm:"not null"`

	Status int `gorm:"default:1"`

	Email string `gorm:"not null"`

	Unit string `gorm:"not null"`

	JoinDate *time.Time

	ExperienceYears int

	CreatedAt time.Time

	UpdatedAt time.Time

	Certifications []K9TrainerCertification `gorm:"foreignKey:TrainerID"`

	// K9TrainingRecords []TrainingRecord.TrainingRecord `gorm:"foreignKey:TrainerID"`
}