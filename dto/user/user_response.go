package user

import (
	"github.com/google/uuid"

	"time"
)

type UserDetailResponse struct {

	ID uuid.UUID `json:"id"`

	FullName string `json:"full_name"`

	Email string `json:"email"`

	Department string `json:"department"`

	Status int `json:"status"`

	RoleName string `json:"role_name"`

	RoleID uuid.UUID `json:"role_id"`

	LastLoginAt *time.Time `json:"last_login_at"`

	AdditionalPermissions []UserPermissionResponse `json:"additional_permissions"`
}
