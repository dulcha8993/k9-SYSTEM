package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RoleID uuid.UUID `gorm:"type:uuid;not null"`
	Role   UserRole  `gorm:"foreignKey:RoleID"`

	FullName string `gorm:"not null"`

	Email string `gorm:"unique;not null"`

	PasswordHash string `gorm:"not null" json:"-"`

	ForcePasswordReset bool `gorm:"default:false"`

	Department string

	Status int `gorm:"default:1"`

	LastLoginAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time

	// AdditionalPermissions []UserAdditionalModuleAccess `gorm:"foreignKey:UserID"` // to keep additional permission done
}