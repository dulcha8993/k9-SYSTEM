package middleware

import (
	"k9-system/config"
	// "k9-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

import (
	accessModel "k9-system/models/access"
)

func PermissionMiddleware(
	moduleKey string,
	action string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		roleIDString := c.GetString("role_id") //Reads User From JWT

		userIDString := c.GetString("user_id")

		roleID, _ := uuid.Parse(roleIDString)

		userID, _ := uuid.Parse(userIDString)

		// Find module
		var module accessModel.SystemModule

		config.DB.
			Where("module_key = ?", moduleKey).
			First(&module)

		if module.ID == uuid.Nil {

			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Module not found",
			})

			c.Abort()

			return
		}

		// ROLE PERMISSION
		var rolePermission accessModel.ModuleAccess

		config.DB.
			Where(
				"role_id = ? AND module_id = ?",
				roleID,
				module.ID,
			).
			First(&rolePermission)

		// USER ADDITIONAL PERMISSION
		var userPermission accessModel.UserAdditionalModuleAccess

		config.DB.
			Where(
				"user_id = ? AND module_id = ?",
				userID,
				module.ID,
			).
			First(&userPermission)

		// MERGE PERMISSIONS
		canAccess := false

		switch action {

		case "view":

			canAccess =
				rolePermission.CanView ||
					userPermission.CanView

		case "create":

			canAccess =
				rolePermission.CanCreate ||
					userPermission.CanCreate

		case "update":

			canAccess =
				rolePermission.CanUpdate ||
					userPermission.CanUpdate

		case "delete":

			canAccess =
				rolePermission.CanDelete ||
					userPermission.CanDelete
		}

		if !canAccess {

			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Permission denied",
			})

			c.Abort()

			return
		}

		c.Next()
	}
}