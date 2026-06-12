package k9_trainer

import (
	"time"

	"github.com/google/uuid"
)		

type K9TrainerCertification struct {

	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"-"`

	TrainerID uuid.UUID `gorm:"type:uuid;not null"`
	// K9Trainer   K9Trainer `gorm:"foreignKey:TrainerID;"`

	Name string `gorm:"not null"`

	IssuingOrganization string `gorm:"default:null"`

	DateObtained *time.Time

	DateExpires *time.Time

	CreatedAt time.Time `json:"-"`

	UpdatedAt time.Time `json:"-"`
	}