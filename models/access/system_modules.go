package access

import (
	"time"

	"github.com/google/uuid"
)

type SystemModule struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name               string
	ModuleKey          string `gorm:"unique;not null"`
	Status             int `gorm:"default:1"`

	CreatedAt          time.Time
	UpdatedAt          time.Time
}