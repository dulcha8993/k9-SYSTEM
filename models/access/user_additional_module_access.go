package access

import (
	"time"

	"github.com/google/uuid"
)

import (
	User "k9-system/models/user"
)

type UserAdditionalModuleAccess struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_module"`
	// User User `gorm:"foreignKey:UserID"`

	UserID uuid.UUID
	User User.User `gorm:"foreignKey:UserID" json:"-"` // json:"-" meaning - do not include this in api response

	ModuleID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_module"`
	Module SystemModule `gorm:"foreignKey:ModuleID"`

	CanView bool `gorm:"default:false"`
	CanCreate bool `gorm:"default:false"`
	CanUpdate bool `gorm:"default:false"`
	CanDelete bool `gorm:"default:false"`

	Status int `gorm:"default:1"`

	CreatedAt time.Time
	UpdatedAt time.Time
}