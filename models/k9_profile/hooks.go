package k9_profile

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *K9Profile) BeforeCreate(tx *gorm.DB) error {

	r.ID = uuid.New()

	return nil
}