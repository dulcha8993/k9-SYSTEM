package user


type UpdateUserRequest struct {

	RoleID *string `json:"role_id"`

	FullName *string `json:"full_name"`

	Email *string `json:"email"`

	Password *string `json:"password"`

	Department *string `json:"department"`

	Status *int `json:"status"`

	AdditionalPermissions []AdditionalPermissionRequest `json:"additional_permissions"`
}