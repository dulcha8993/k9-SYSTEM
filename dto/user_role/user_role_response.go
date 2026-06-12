package user_role

import (
	"github.com/google/uuid"
)

type ModuleAccessResponse struct {

	ID uuid.UUID `json:"id"`

	ModuleID uuid.UUID `json:"module_id"`

	ModuleName string `json:"module_name"`

	ModuleKey string `json:"module_key"`

	CanView bool `json:"can_view"`

	CanCreate bool `json:"can_create"`

	CanUpdate bool `json:"can_update"`

	CanDelete bool `json:"can_delete"`

	Status int `json:"status"`
}

type UserRoleResponse struct {

	ID uuid.UUID `json:"id"`

	RoleName string `json:"role_name"`

	Status int `json:"status"`

	ModuleAccess []ModuleAccessResponse `json:"module_access"`
}