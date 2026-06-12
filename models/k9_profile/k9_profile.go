package k9_profile

import (
	"time"

	"github.com/google/uuid"
)

// import (
// 	TrainingRecord "k9-system/models/training_record"
// )

type K9Profile struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	K9Code string `gorm:"unique;not null"` //PostgreSQL sequence and transaction lock and serial counter table

	Name string `gorm:"not null"`

	Breed string `gorm:"not null"`

	Gender string `gorm:"not null"`

	DateOfBirth *time.Time

	Microchip string `gorm:"not null"`

	Status int `gorm:"default:1"`

	Color string `gorm:"not null"`

	ServiceSince *time.Time

	ServiceEnd *time.Time

	CreatedAt time.Time

	UpdatedAt time.Time

	// K9TrainingRecords []TrainingRecord.TrainingRecord `gorm:"foreignKey:K9ID"`
}