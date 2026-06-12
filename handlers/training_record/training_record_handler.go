package training_record

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

import (
	// config "k9-system/config"
	// TrainingRecord "k9-system/models/training_record"
	utils "k9-system/utils"
)

import (
	dto "k9-system/dto/training_record"
	response "k9-system/response"
	trainingRecordService "k9-system/services/training_record"
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

	trainingRecord, err := trainingRecordService.UpdateTrainingRecord(
		id,
		req,
	)

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