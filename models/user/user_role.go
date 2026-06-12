package user

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleName  string    `gorm:"not null"`
	Status    int `gorm:"default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// ModuleAccess []ModuleAccess `gorm:"foreignKey:RoleID"` // to keep additional permission done
}