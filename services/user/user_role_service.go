package services

import (
	"k9-system/config"
	"errors"
	// "github.com/gin-gonic/gin"
)

import (
	userModel "k9-system/models/user"
	accessModel "k9-system/models/access"
	dto "k9-system/dto/user_role"
	constants "k9-system/constants"
)

func GetUserRoles() (
	[]dto.UserRoleResponse,
	error,
) {

	var userRoles []userModel.UserRole

	if err := config.DB.
		Find(&userRoles).
		Error; err != nil {

		return nil, errors.New(
			constants.FailedToRetrieve(
				"user roles",
			),
		)
	}

	var userRoleResult []dto.UserRoleResponse

	for _, userRole := range userRoles {

		var accesses []accessModel.ModuleAccess

		if err := config.DB.
			Preload("SystemModule").
			Where(
				"role_id = ?",
				userRole.ID,
			).
			Find(&accesses).
			Error; err != nil {

			return nil, errors.New(
				constants.FailedToRetrieve(
					"module access",
				),
			)
		}

		var moduleAccess []dto.ModuleAccessResponse

		for _, access := range accesses {

			moduleAccess = append(
				moduleAccess,
				dto.ModuleAccessResponse{

					ID: access.ID,

					ModuleID: access.ModuleID,

					ModuleName: access.SystemModule.Name,

					ModuleKey: access.SystemModule.ModuleKey,

					CanView: access.CanView,

					CanCreate: access.CanCreate,

					CanUpdate: access.CanUpdate,

					CanDelete: access.CanDelete,

					Status: access.Status,
				},
			)
		}

		roleResponse := dto.UserRoleResponse{

			ID: userRole.ID,

			RoleName: userRole.RoleName,

			Status: userRole.Status,

			ModuleAccess: moduleAccess,
		}

		userRoleResult = append(
			userRoleResult,
			roleResponse,
		)
	}

	return userRoleResult, nil
}