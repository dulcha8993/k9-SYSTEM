package database

import (
	"k9-system/config"
	// "k9-system/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

import (
	userModel "k9-system/models/user"
)

func SeedAdminUser() {

	var existingUser userModel.User
	config.DB.
		Where("email = ?", "admin@k9.com").
		First(&existingUser)

	if existingUser.ID != uuid.Nil {
		return
	}

	// Get admin role
	var adminRole userModel.UserRole

	config.DB.
		Where("role_name = ?", "System admin").
		First(&adminRole)

	if adminRole.ID == uuid.Nil {
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("123456"),
		bcrypt.DefaultCost,
	)

	user := userModel.User{
		RoleID: adminRole.ID,

		FullName: "Test System Admin",

		Email: "admin@k9.com",

		PasswordHash: string(hashedPassword),

		Department: "System Administration",

		ForcePasswordReset: false,

		Status: 1,
	}

	config.DB.Create(&user)
}