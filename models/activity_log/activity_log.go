package activity_log

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/datatypes"
)

type ActivityLog struct {

	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`

	Module string `gorm:"size:100;not null"`

	Action string `gorm:"size:50;not null"`

	RecordID *uuid.UUID `gorm:"type:uuid;not null"`

	Description string `gorm:"type:text"`

	ErrorMessage *string `gorm:"type:text"`

	Changes datatypes.JSON `gorm:"type:jsonb"`

	Status string `gorm:"type:text"`

	CreatedAt time.Time
}