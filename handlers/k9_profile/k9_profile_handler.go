package k9_profile

import (
	"net/http"
	// "fmt"
	"gorm.io/gorm" 

	"github.com/gin-gonic/gin"
)

import (
	utils "k9-system/utils"
	"github.com/google/uuid"
)

import (
	dto "k9-system/dto/k9_profile"
	response "k9-system/response"
	k9ProfileService "k9-system/services/k9_profile"
	activityLogService "k9-system/services/activity_log"
	activityLogDTO "k9-system/dto/activity_log"
	"k9-system/constants"
)

func CreateK9Profile(c *gin.Context) {

	var req dto.CreateK9ProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	record, err := k9ProfileService.CreateK9Profile(
		req,
	)

	LogCreate(c, constants.ActionCreate ,&record.ID, err, nil)

	if err != nil {

		response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	response.Created(c,record,)


}

func UpdateK9Profile(c *gin.Context) {

	var req dto.UpdateK9ProfileRequest

	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	record, changes, err := k9ProfileService.UpdateK9Profile(
		id,
		req,
	)

	LogCreate(c, constants.ActionUpdate ,&record.ID, err, changes)

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
		record,
	)
}

func GetK9Profiles(c *gin.Context) {

	search := c.Query("search")

	status := c.Query("status")

	microchip := c.Query("microchip")	

	pagination := utils.GetPagination(c)

	
	k9Profiles, pagination, err := k9ProfileService.GetK9Profiles(
		search,
		status,
		microchip,
		pagination,
	)
	LogCreate(c, constants.ActionViewList ,nil, err, nil)

	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	response.Paginated(
		c,
		k9Profiles,
		pagination,
	)
}

func GetK9ProfileByID(c *gin.Context) {

	id := c.Param("id")

	
	k9Profile, err :=  k9ProfileService.GetK9ProfileByID(
		id,
	)

	LogCreate(c, constants.ActionView ,&k9Profile.ID, err, nil)

	if err != nil {

		if err == gorm.ErrRecordNotFound {		
			response.Error(
				c,
				http.StatusNotFound,
				err.Error(),
			)
		} else {
			response.Error(
				c,
				http.StatusInternalServerError,
				err.Error(),
			)
		}
		return
	}
	response.Success(c,k9Profile,)
}

func LogCreate(
	c *gin.Context,
	action string,
	recordID *uuid.UUID,
	err error,
	changes []activityLogDTO.FieldChange,
) {

	userID := c.MustGet("user_id").(uuid.UUID)

	description := action + " K9 Profile"

	if recordID != nil {
		description += ": " + recordID.String()
	}

	if err != nil {

		activityLogService.LogFailure(
			userID,
			"k9_profile",
			action,
			description,
			err,
		)

		return
	}

	activityLogService.LogSuccess(
		userID,
		"k9_profile",
		action,
		recordID,
		description,
		changes,
	)
}