package user

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin" // HTTP web framework for Go (Golang)
)

import (
	utils "k9-system/utils"
	"github.com/google/uuid"
)

import (
	dto "k9-system/dto/user"
	response "k9-system/response"
	userService "k9-system/services/user"
	activityLogService "k9-system/services/activity_log"
	activityLogDTO "k9-system/dto/activity_log"
	"k9-system/constants"
)

func CreateUser(c *gin.Context) {

	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	user, err := userService.CreateUser(
		req,
	)

	if err != nil {

		LogCreate(
			c,
			constants.ActionCreate,
			nil,
			err,
			nil,
		)

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	LogCreate(
		c,
		constants.ActionCreate,
		&user.ID,
		nil,
		nil,
	)

	response.Created(
		c,
		user,
	)
}

func UpdateUser(c *gin.Context) {

	var req dto.UpdateUserRequest

	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	user, changes, err := userService.UpdateUser(
		id,
		req,
	)

	LogCreate(c, constants.ActionUpdate ,&user.ID, err, changes)

	if err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	response.Success(
		c,
		user,
	)
}

func GetUsers(c *gin.Context) {

	search := c.Query("search")

	status := c.Query("status")

	pagination := utils.GetPagination(c)

	users, pagination, err := userService.GetUsers(
		search,
		status,
		pagination,
	)

		LogCreate(c, constants.ActionViewList ,nil, err, nil)

	if err != nil {

		response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	response.Paginated(
		c,
		users,
		pagination,
	)
}


func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	userResponse, err := userService.GetUserByID(
		id,
	)

	LogCreate(c, constants.ActionView ,&userResponse.ID, err, nil)

	if err != nil {

		response.Error(
			c,
			http.StatusNotFound,
			err.Error(),
		)

		return
	}

	response.Success(
		c,
		userResponse,
	)
}

func LogCreate(
	c *gin.Context,
	action string,
	recordID *uuid.UUID,
	err error,
	changes []activityLogDTO.FieldChange,
) {

	userID := c.MustGet("user_id").(uuid.UUID)

	description := action + " System admin"

	if recordID != nil {
		description += ": " + recordID.String()
	}

	if err != nil {

		activityLogService.LogFailure(
			userID,
			"System_admin",
			action,
			description,
			err,
		)

		return
	}

	activityLogService.LogSuccess(
		userID,
		"System_admin",
		action,
		recordID,
		description,
		changes,
	)
}