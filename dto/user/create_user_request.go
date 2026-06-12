package user

type CreateUserRequest struct {
	RoleID string `json:"role_id" binding:"required"`

	FullName string `json:"full_name" binding:"required"`

	Email string `json:"email" binding:"required,email"`

	Password string `json:"password" binding:"required,min=6"`

	Department string `json:"department"`

	ForcePasswordReset bool `json:"force_password_reset"`

	Status int `json:"status" binding:"required"`

	AdditionalPermissions []AdditionalPermissionRequest `json:"additional_permissions"`
}