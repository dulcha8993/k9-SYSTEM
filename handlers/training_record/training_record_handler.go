package training_record

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

import (
	utils "k9-system/utils"
	"github.com/google/uuid"
)

import (
	dto "k9-system/dto/training_record"
	response "k9-system/response"
	trainingRecordService "k9-system/services/training_record"
	activityLogService "k9-system/services/activity_log"
	activityLogDTO "k9-system/dto/activity_log"
	"k9-system/constants"
)

func CreateTraningRecord(c *gin.Context) {

	var req dto.CreateTrainingRecordRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	trainingRecord, err := trainingRecordService.CreateTrainingRecord(
		req,
	)

	LogCreate(c, constants.ActionCreate ,&trainingRecord.ID, err, nil)

	if err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	response.Created(
		c,
		trainingRecord,
	)
}

func UpdateTrainingRecord(c *gin.Context) {

	var req dto.UpdateTrainingRecordRequest

	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	trainingRecord, changes, err := trainingRecordService.UpdateTrainingRecord(
		id,
		req,
	)

	LogCreate(c, constants.ActionUpdate ,&trainingRecord.ID, err, changes)

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
		trainingRecord,
	)
}

func GetTrainingRecords(c *gin.Context) {

	search := c.Query("search")

	status := c.Query("status")

	trainingRecordID := c.Query("training_record_id")

	pagination := utils.GetPagination(c)

	trainingRecords, pagination, err := trainingRecordService.GetTrainingRecords(
		search,
		status,
		trainingRecordID,
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
		trainingRecords,
		pagination,
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

	description := action + " K9 Training"

	if recordID != nil {
		description += ": " + recordID.String()
	}

	if err != nil {

		activityLogService.LogFailure(
			userID,
			"k9_training",
			action,
			description,
			err,
		)

		return
	}

	activityLogService.LogSuccess(
		userID,
		"k9_training",
		action,
		recordID,
		description,
		changes,
	)
}