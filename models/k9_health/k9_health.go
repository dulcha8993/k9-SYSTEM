package k9_health

import (
	"time"

	"github.com/google/uuid"
	K9ProfileModel "k9-system/models/k9_profile"
)

type HealthRecord struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	K9ID uuid.UUID `gorm:"type:uuid;not null"`
	K9Profile   K9ProfileModel.K9Profile `gorm:"foreignKey:K9ID;"`

	VisitedDate *time.Time

	ReleasedDate *time.Time

	ClinicName string `gorm:"not null"`

	Diagnosis string `gorm:"not null"`
	
	NextAppointment *time.Time

	Veterinarian string  `gorm:"not null"`

	Treatment string `gorm:"not null"`

	Type string `gorm:"not null"`

	Medication string `gorm:"not null"`

	Description string

	RecoveryStatus string

	CreatedAt time.Time

	UpdatedAt time.Time
}