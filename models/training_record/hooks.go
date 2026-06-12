package training_record

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *TrainingRecord) BeforeCreate(tx *gorm.DB) error {

	r.ID = uuid.New()

	return nil
}