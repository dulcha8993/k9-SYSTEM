
package k9_health

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *HealthRecord) BeforeCreate(tx *gorm.DB) error {

	r.ID = uuid.New()

	return nil
}