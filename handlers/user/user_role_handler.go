package user

import (
	// "k9-system/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

import (
	// dto "k9-system/dto/user_role"
	response "k9-system/response"
	userService "k9-system/services/user" 
)



func GetUserRoles(c *gin.Context) {

	userRoles, err := userService.GetUserRoles()

	if err != nil {

		response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	response.Success(
		c,
		userRoles,
	)
}