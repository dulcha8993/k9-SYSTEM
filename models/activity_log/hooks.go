
package activity_log

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *ActivityLog) BeforeCreate(tx *gorm.DB) error {

	r.ID = uuid.New()

	return nil
}