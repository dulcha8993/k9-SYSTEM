package k9_trainer

import (
	"time"

	"github.com/google/uuid"
)

import ( 
	K9Profile "k9-system/models/k9_profile"
)

type K9TrainerHistory struct {

	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	TrainerID uuid.UUID `gorm:"type:uuid;not null"`
	Trainer K9Trainer `gorm:"foreignKey:TrainerID;"`

	K9ID *uuid.UUID `gorm:"type:uuid;default:null"`
	K9Profile K9Profile.K9Profile `gorm:"foreignKey:K9ID;"`

	StartDate *time.Time

	EndDate *time.Time

	Status int `gorm:"default:1"`

	CreatedAt time.Time `json:"-"`
}	