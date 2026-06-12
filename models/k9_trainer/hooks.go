package k9_trainer

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *K9Trainer) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}

func (r *K9TrainerCertification) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}

func (r *K9TrainerHistory) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}