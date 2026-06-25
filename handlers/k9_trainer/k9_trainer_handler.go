package k9_trainer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

import (
	utils "k9-system/utils"
	"github.com/google/uuid"
)

import (
	response "k9-system/response"
	dto "k9-system/dto/k9_trainer"
	k9TrainerService "k9-system/services/k9_trainer"
	activityLogService "k9-system/services/activity_log"
	activityLogDTO "k9-system/dto/activity_log"
	"k9-system/constants"
)

func CreateK9Trainer(c *gin.Context) {
		var req dto.K9TrainerRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			response.Error(
					c,
					http.StatusBadRequest,
					err.Error(),
			)
			return
		}

	trainerResponse, err := k9TrainerService.CreateK9Trainer(
		req,
	)

	LogCreate(c, constants.ActionCreate ,&trainerResponse.GenaralInfo.ID, err, nil)

	if err != nil {

		response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	response.Created(c,trainerResponse,)
}

func UpdateK9Trainer(c *gin.Context) {

	var req dto.K9TrainerRequest
	// var req dto.K9TrainerUpdateRequest

	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	trainer, changes, err := k9TrainerService.UpdateK9Trainer(
		id,
		req,
	)

	LogCreate(c, constants.ActionUpdate ,&trainer.ID, err, changes)

	if err != nil {

		response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	response.Success(c,trainer,)
}

func GetK9Trainers(c *gin.Context) {

	search := c.Query("search")

	status := c.Query("status")

	trainer_id := c.Query("trainer_id")

	pagination := utils.GetPagination(c)

	trainers, pagination, err := k9TrainerService.GetK9Trainers(
		search,
		status,
		trainer_id,
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
		trainers,
		pagination,
	)
}

func AssignK9TrainerToK9(c *gin.Context) {

	var req dto.CreateK9TrainerAssignRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	result, err := k9TrainerService.AssignK9TrainerToK9(
		req,
	)

	LogCreate(c, constants.ActionAssignK9 ,&result.ID, err, nil)

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
		result,
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

	description := action + " K9 Trainer"

	if recordID != nil {
		description += ": " + recordID.String()
	}

	if err != nil {

		activityLogService.LogFailure(
			userID,
			"k9_trainer",
			action,
			description,
			err,
		)

		return
	}

	activityLogService.LogSuccess(
		userID,
		"k9_trainer",
		action,
		recordID,
		description,
		changes,
	)
}