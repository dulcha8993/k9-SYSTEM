package training_record

import (
	"time"

	"github.com/google/uuid"
)

import (
	K9TrainerModel "k9-system/models/k9_trainer"
	K9ProfileModel "k9-system/models/k9_profile"

)

type TrainingRecord struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	K9ID uuid.UUID `gorm:"type:uuid;not null"`
	K9Profile   K9ProfileModel.K9Profile `gorm:"foreignKey:K9ID;"`

	TrainerID uuid.UUID `gorm:"type:uuid;not null"`
	K9Trainer   K9TrainerModel.K9Trainer `gorm:"foreignKey:TrainerID;"`

	TrainerName string `gorm:"not null"`

	TrainingType string `gorm:"not null"`

	TrainingModule string `gorm:"not null"`

	StartDate time.Time `gorm:"not null"`

	EndDate time.Time `gorm:"not null"`

	TrainingCenterName string `gorm:"not null"`

	TrainingCenterAddressLine1 string `gorm:"not null"`

	TrainingCenterAddressLine2 string 

	TrainingCenterCountry string `gorm:"not null"`

	PerformanceScore float64 `gorm:"not null"`

	Notes string `gorm:"type:text"`

	Objectives string `gorm:"type:text"`

	CreatedAt time.Time 

	UpdatedAt time.Time

}