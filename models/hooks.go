package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// u *User - method reciver - this function belongs to user model - change the origina uuid value of the user without creating a copy

// if method modify values use pointers:  func (u *User)
//if method only read values user method reciver: func (u User)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

func (r *UserRole) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}

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

func (r *K9Profile) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}

func (r *K9Trainer) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}


// CREATE TYPE access_level AS ENUM (
//   'NO_ACCESS',
//   'VIEW',
//   'CREATE_VIEW',
//   'VIEW_UPDATE',
//   'CREATE_UPDATE_VIEW',
//   'FULL',
//   'VIEW_DELETE',
//   'CREATE_DELETE',
//   'UPDATE_DELETE',
//   'CREATE_UPDATE_DELETE',
// );