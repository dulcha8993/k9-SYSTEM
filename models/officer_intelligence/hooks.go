package officer_intelligence

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *CriminalCase) BeforeCreate(tx *gorm.DB) error {

	r.ID = uuid.New()

	return nil
}