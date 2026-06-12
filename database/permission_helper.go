package database

import (
	"k9-system/config"
	// "k9-system/models"

	"github.com/google/uuid"
)

import (
	userModel "k9-system/models/user"
	accessModel "k9-system/models/access"
)

func CreatePermission(
	roleName string,
	moduleKey string,
	canView bool,
	canCreate bool,
	canUpdate bool,
	canDelete bool,
) {

	var role userModel.UserRole

	config.DB.
		Where("role_name = ?", roleName).
		First(&role)

	if role.ID == uuid.Nil {
		return
	}

	var module accessModel.SystemModule

	config.DB.
		Where("module_key = ?", moduleKey).
		First(&module)

	if module.ID == uuid.Nil {
		return
	}

	var existing accessModel.ModuleAccess

	config.DB.
		Where("role_id = ? AND module_id = ?",
			role.ID,
			module.ID).
		First(&existing)

	if existing.ID != uuid.Nil {
		return
	}

	permission := accessModel.ModuleAccess{
		RoleID: role.ID,
		ModuleID: module.ID,

		CanView: canView,
		CanCreate: canCreate,
		CanUpdate: canUpdate,
		CanDelete: canDelete,

		Status: 1,
	}

	config.DB.Create(&permission)
}