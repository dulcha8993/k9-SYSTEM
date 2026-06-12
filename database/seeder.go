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


func SeedRoles() {

	roles := []userModel.UserRole{
		{RoleName: "System admin"},
		{RoleName: "K9 operation officer"},
		{RoleName: "K9 handler"},
		{RoleName: "K9 veterinary"},
		{RoleName: "K9 trainer"},
		{RoleName: "Officer intelligence"},
	}

	for _, role := range roles {

		var existingRole userModel.UserRole

		config.DB.
			Where("role_name = ?", role.RoleName).
			First(&existingRole)

		if existingRole.ID == uuid.Nil {

			config.DB.Create(&role)
		}
	}
}

func SeedSystemModules() {

	modules := []accessModel.SystemModule{
		{
			Name:      "System dmin - Governance, UAM, Retirement",
			ModuleKey: "System_admin",
		},
		{
			Name:      "k9 Profile Management - K9 Pfofile Management",
			ModuleKey: "K9_profile_management",
		},
		{
			Name:      "k9 Handler - Daily Opertations",
			ModuleKey: "k9_handler",
		},
		{
			Name:      "K9 Veterianary - Health and Medical Records",
			ModuleKey: "k9_veterianary",
		},
		{
			Name:      "K9 Trainer - Training & Certification",
			ModuleKey: "k9_trainer",
		},
		{
			Name:      "Officer Intelligence - Criminal Case Management",
			ModuleKey: "officer_intelligence",
		},
	}

	for _, module := range modules {

		var existingModule accessModel.SystemModule

		config.DB.
			Where("module_key = ?", module.ModuleKey).
			First(&existingModule)

		if existingModule.ID == uuid.Nil {

			config.DB.Create(&module)
		}
	}
}

func SeedModuleAccess() {

	CreatePermission(
		"System admin",
		"System_admin",
		true,
		true,
		true,
		true,
	)

	CreatePermission(
		"System admin",
		"K9_profile_management",
		true,
		true,
		true,
		true,
	)

	CreatePermission(
		"System admin",
		"k9_handler",
		true,
		true,
		true,
		true,
	)

	CreatePermission(
		"System admin",
		"k9_training",
		true,
		true,
		true,
		true,
	)

	CreatePermission(
		"System admin",
		"k9_veterianary",
		true,
		true,
		true,
		true,
	)

	CreatePermission(
		"System admin",
		"officer_intelligence",
		true,
		true,
		true,
		false,
	)

	// K9 TRAINER
	CreatePermission(
		"K9 trainer",
		"k9_trainer",
		true,
		true,
		true,
		false,
	)

	// K9 HANDLER
	CreatePermission(
		"K9 handler",
		"k9_handler",
		true,
		false,
		false,
		false,
	)

	// VETERINARY
	CreatePermission(
		"K9 veterinary",
		"k9_veterianary",
		true,
		true,
		true,
		false,
	)

	// OFFICER INTEL
	CreatePermission(
		"Officer intelligence",
		"officer_intelligence",
		true,
		true,
		true,
		false,
	)
}
