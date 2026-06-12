package officer_intelligence

import (
	"time"

	"github.com/google/uuid"
)

import (
	K9TrainerHistory "k9-system/models/k9_trainer"
)

type CriminalCase struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	K9TrainerHistoryID uuid.UUID `gorm:"type:uuid;not null"`
	TrainerHistory   K9TrainerHistory.K9TrainerHistory `gorm:"foreignKey:K9TrainerHistoryID;"`

	CaseNumber string `gorm:"not null"`

	Title string `gorm:"not null"`

	Description string `gorm:"type:text"`

	Outcome string `gorm:"type:text"`

	Status int `gorm:"not null"`

	OpenDate time.Time `gorm:"not null"`

	ClosedDate *time.Time

	Role string `gorm:"not null"`

	Location string `gorm:"not null"`

	ReportingOfficer string `gorm:"not null"`

	CreatedAt time.Time

	UpdatedAt time.Time
}