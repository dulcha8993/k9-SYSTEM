package user

import (
	"github.com/google/uuid"
)

type AdditionalPermissionRequest struct {
	ModuleID string `json:"module_id"`

	CanView bool `json:"can_view"`
	CanCreate bool `json:"can_create"`
	CanUpdate bool `json:"can_update"`
	CanDelete bool `json:"can_delete"`
}

type UserPermissionResponse struct {

	ID uuid.UUID `json:"id"`

	ModuleID uuid.UUID `json:"module_id"`

	ModuleName string `json:"module_name"`

	ModuleKey string `json:"module_key"`

	CanView bool `json:"can_view"`

	CanCreate bool `json:"can_create"`

	CanUpdate bool `json:"can_update"`

	CanDelete bool `json:"can_delete"`
}