package access

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *SystemModule) BeforeCreate(tx *gorm.DB) error {

	r.ID = uuid.New()

	return nil
}

func (r *ModuleAccess) BeforeCreate(tx *gorm.DB) error {

	r.ID = uuid.New()

	return nil
}

func (r *UserAdditionalModuleAccess) BeforeCreate(tx *gorm.DB) error {

	r.ID = uuid.New()

	return nil
}