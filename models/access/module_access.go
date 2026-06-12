package access

import (
	"time"

	"github.com/google/uuid"
)

import (
	UserRole "k9-system/models/user"
)

type ModuleAccess struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RoleID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_role_module"`
	Role   UserRole.UserRole  `gorm:"foreignKey:RoleID"`

	ModuleID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_role_module"`
	SystemModule SystemModule `gorm:"foreignKey:ModuleID"`

	CanCreate bool `gorm:"default:false"`
	CanUpdate bool `gorm:"default:false"`
	CanView   bool `gorm:"default:false"`
	CanDelete bool `gorm:"default:false"`

	Status int `gorm:"default:1"`

	CreatedAt time.Time
	UpdatedAt time.Time
}