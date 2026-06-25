package k9_health

import (
	"net/http"
	"gorm.io/gorm" 
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

import (
	utils "k9-system/utils"
)

import (
	dto "k9-system/dto/k9_health"
	response "k9-system/response"
	k9HealthService "k9-system/services/k9_health"
	activityLogService "k9-system/services/activity_log"
	activityLogDTO "k9-system/dto/activity_log"
	"k9-system/constants"
)


func CretateHealthRecord(c *gin.Context) {

	var req dto.CreateK9HealthRecordRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	record, err := k9HealthService.CreateHealthRecord(
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

	response.Created(
		c,
		record,
	)
}

func UpdateHealthRecord(c *gin.Context) {

	var req dto.UpdateK9HealthRecordRequest

	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),	
		)
		return
	}

	record, changes, err := k9HealthService.UpdateHealthRecord(
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

func GetHealthRecords(c *gin.Context) {

	search := c.Query("search")
	status := c.Query("status")
	healthRecordType := c.Query("type")

	pagination := utils.GetPagination(c)

	healthRecords, pagination, err := k9HealthService.GetHealthRecords(
		search,
		status,
		healthRecordType,
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
		healthRecords,
		pagination,
	)
}

func GetHealthRecordByID(c *gin.Context) {

	id := c.Param("id");

	K9HealthRecord, err :=  k9HealthService.GetaHealthRecordByID(
		id,
	)

	LogCreate(c, constants.ActionView ,&K9HealthRecord.ID, err, nil)

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

	response.Success(c,K9HealthRecord,)
}

func LogCreate(
	c *gin.Context,
	action string,
	recordID *uuid.UUID,
	err error,
	changes []activityLogDTO.FieldChange,
) {

	userID := c.MustGet("user_id").(uuid.UUID)

	description := action + " K9 Health Record"

	if recordID != nil {
		description += ": " + recordID.String()
	}

	if err != nil {

		activityLogService.LogFailure(
			userID,
			"k9_veterianary",
			action,
			description,
			err,
		)

		return
	}

	activityLogService.LogSuccess(
		userID,
		"k9_veterianary",
		action,
		recordID,
		description,
		changes,
	)
}