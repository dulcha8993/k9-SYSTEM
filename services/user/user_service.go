package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"

	config "k9-system/config"

	"k9-system/constants"

	dto "k9-system/dto/user"

	"k9-system/utils"

	accessModel "k9-system/models/access"

	userModel "k9-system/models/user"

	activityLogDTO "k9-system/dto/activity_log"
)

func CreateUser(
	req dto.CreateUserRequest,
) (*userModel.User, error) {

	// Check existing email

	var existingUser userModel.User

	config.DB.
		Where("email = ?", req.Email).
		First(&existingUser)

	if existingUser.ID != uuid.Nil {

		return nil, errors.New(
			"Email already exists",
		)
	}

	// Parse role UUID

	roleUUID, err := utils.ParseUUID(
		req.RoleID,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidUUID(
				"User Role ID",
			),
		)
	}

	// Hash password

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {

		return nil, errors.New(
			"Password hashing failed",
		)
	}

	tx := config.DB.Begin()

	user := userModel.User{

		RoleID: roleUUID,

		FullName: req.FullName,

		Email: req.Email,

		PasswordHash: string(hashedPassword),

		Department: req.Department,

		ForcePasswordReset: req.ForcePasswordReset,

		Status: req.Status,
	}

	if err := tx.Create(&user).Error; err != nil {

		tx.Rollback()

		return nil, err
	}

	// Create additional permissions

	for _, permission := range req.AdditionalPermissions {

		moduleUUID, err := utils.ParseUUID(
			permission.ModuleID,
		)

		if err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.InvalidUUID(
					"User Module ID",
				),
			)
		}

		additionalPermission := accessModel.UserAdditionalModuleAccess{

			UserID: user.ID,

			ModuleID: moduleUUID,

			CanView: permission.CanView,

			CanCreate: permission.CanCreate,

			CanUpdate: permission.CanUpdate,

			CanDelete: permission.CanDelete,

			Status: 1,
		}

		if err := tx.
			Create(&additionalPermission).
			Error; err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.FailedToCreate(
					"user permissions",
				),
			)
		}
	}

	// if err := tx.
	// 	Preload("Role").
	// 	Preload("AdditionalPermissions").
	// 	First(&user, "id = ?", user.ID).
	// 	Error; err != nil {

	// 	tx.Rollback()

	// 	return nil, nil, errors.New(
	// 		constants.FailedToRetrieve(
	// 			"user",
	// 		),
	// 	)
	// }

	if err := tx.Commit().Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func UpdateUser(
	id string,
	req dto.UpdateUserRequest,
) (*userModel.User, []activityLogDTO.FieldChange, error) {

	var existingUser userModel.User

	if err := config.DB.
		Where("id = ?", id).
		First(&existingUser).
		Error; err != nil {

		return nil, nil, errors.New(
			constants.NotFound("User"),
		)
	}

	// Keep original copy
	original := existingUser
	// Check duplicate email

	if req.Email != nil {

		var emailUser userModel.User

		config.DB.
			Where(
				"email = ? AND id != ?",
				*req.Email,
				id,
			).
			First(&emailUser)

		if emailUser.ID != uuid.Nil {

			return nil, nil, errors.New(
				"Email already exists",
			)
		}
	}

	tx := config.DB.Begin()

	// Update role

	if req.RoleID != nil {

		roleUUID, err := utils.ParseUUID(
			*req.RoleID,
		)

		if err != nil {

			tx.Rollback()

			return nil, nil, errors.New(
				constants.InvalidUUID(
					"User Role ID",
				),
			)
		}

		var role userModel.UserRole

		tx.
			Where("id = ?", roleUUID).
			First(&role)

		if role.ID == uuid.Nil {

			tx.Rollback()

			return nil, nil, errors.New(
				constants.NotFound(
					"User role",
				),
			)
		}

		existingUser.RoleID = roleUUID
	}

	// Partial Updates

	if req.FullName != nil {
		existingUser.FullName = *req.FullName
	}

	if req.Email != nil {
		existingUser.Email = *req.Email
	}

	if req.Department != nil {
		existingUser.Department = *req.Department
	}

	if req.Status != nil {
		existingUser.Status = *req.Status
	}

	// Password

	if req.Password != nil &&
		*req.Password != "" {

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(*req.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {

			tx.Rollback()

			return nil, nil, errors.New(
				"Password hashing failed",
			)
		}

		existingUser.PasswordHash = string(
			hashedPassword,
		)
	}

	// Detect changes BEFORE Save
	changes := utils.GetChanges(
		original,
		existingUser,
	)


	// Save User
	if err := tx.
		Save(&existingUser).
		Error; err != nil {

		tx.Rollback()

		return nil, nil, errors.New(
			constants.FailedToUpdate(
				"user",
			),
		)
	}

	// Replace Additional Permissions

	if req.AdditionalPermissions != nil {

		if err := tx.
			Where(
				"user_id = ?",
				existingUser.ID,
			).
			Delete(
				&accessModel.UserAdditionalModuleAccess{},
			).
			Error; err != nil {

			tx.Rollback()

			return nil, nil, err
		}

		for _, permission := range req.AdditionalPermissions {

			moduleUUID, err := utils.ParseUUID(
				permission.ModuleID,
			)

			if err != nil {

				tx.Rollback()

				return nil, nil, errors.New(
					constants.InvalidUUID(
						"Module ID",
					),
				)
			}

			var module accessModel.SystemModule

			tx.
				Where(
					"id = ?",
					moduleUUID,
				).
				First(&module)

			if module.ID == uuid.Nil {

				tx.Rollback()

				return nil, nil, errors.New(
					constants.NotFound(
						"Module",
					),
				)
			}

			additionalPermission :=
				accessModel.UserAdditionalModuleAccess{

					UserID: existingUser.ID,

					ModuleID: moduleUUID,

					CanView: permission.CanView,

					CanCreate: permission.CanCreate,

					CanUpdate: permission.CanUpdate,

					CanDelete: permission.CanDelete,

					Status: 1,
				}

			if err := tx.
				Create(
					&additionalPermission,
				).
				Error; err != nil {

				tx.Rollback()

				return nil, nil, errors.New(
					constants.FailedToCreate(
						"user permission",
					),
				)
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, nil, err
	}

	return &existingUser, changes, nil
}

func GetUsers(
	search string,
	status string,
	pagination utils.Pagination,
) ([]userModel.User, utils.Pagination, error) {

	var users []userModel.User

	var total int64

	query := config.DB.Model(
		&userModel.User{},
	)

	// Search filter
	if search != "" {

		searchPattern := "%" + search + "%"

		query = query.Where(
			`
			full_name ILIKE ?
			OR email ILIKE ?
			OR department ILIKE ?
			`,
			searchPattern,
			searchPattern,
			searchPattern,
		)
	}

	// Status filter
	if status != "" {

		query = query.Where(
			"status = ?",
			status,
		)
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {

		return nil, pagination, errors.New(
			constants.FailedToRetrieve(
				"users",
			),
		)
	}

	// Execute query
	if err := query.
		Preload("Role").
		Order("created_at DESC").
		Offset(pagination.Offset).
		Limit(pagination.Limit).
		Find(&users).
		Error; err != nil {

		return nil, pagination, errors.New(
			constants.FailedToRetrieve(
				"users",
			),
		)
	}

	pagination = utils.BuildPagination(
		pagination,
		total,
	)

	return users, pagination, nil
}

func GetUserByID(
	id string,
) (*dto.UserDetailResponse, error) {

	var user userModel.User

	if err := config.DB.
		Preload("Role").
		Where("id = ?", id).
		First(&user).
		Error; err != nil {

		return nil, errors.New(
			constants.NotFound("User"),
		)
	}

	var additionalPermissions []accessModel.UserAdditionalModuleAccess

	if err := config.DB.
		Preload("Module").
		Where("user_id = ?", user.ID).
		Find(&additionalPermissions).
		Error; err != nil {

		return nil, errors.New(
			constants.FailedToRetrieve(
				"user permissions",
			),
		)
	}

	var permissions []dto.UserPermissionResponse

	for _, permission := range additionalPermissions {

		permissions = append(
			permissions,
			dto.UserPermissionResponse{
				ID:         permission.ID,
				ModuleID:   permission.ModuleID,
				ModuleName: permission.Module.Name,
				ModuleKey:  permission.Module.ModuleKey,
				CanView:    permission.CanView,
				CanCreate:  permission.CanCreate,
				CanUpdate:  permission.CanUpdate,
				CanDelete:  permission.CanDelete,
			},
		)
	}

	userResponse := dto.UserDetailResponse{
		ID:                    user.ID,
		FullName:              user.FullName,
		Email:                 user.Email,
		Department:            user.Department,
		Status:                user.Status,
		RoleID:                user.Role.ID,
		LastLoginAt:           user.LastLoginAt,
		RoleName:              user.Role.RoleName,
		AdditionalPermissions: permissions,
	}

	return &userResponse, nil
}

