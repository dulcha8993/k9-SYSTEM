package officer_intelligence

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/google/uuid"
	"k9-system/utils"
)

import (
	dto "k9-system/dto/officer_intelligence"
	response "k9-system/response"
	constants "k9-system/constants"
	criminalCaseService "k9-system/services/officer_intelligence"
	activityLogService "k9-system/services/activity_log"
	activityLogDTO "k9-system/dto/activity_log"
)

func CreateCriminalCase(c *gin.Context) {

	var req dto.CreateCriminalCaseRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	criminalCase, err := criminalCaseService.CreateCriminalCase(
		req,
	)

	LogCreate(c, constants.ActionCreate ,&criminalCase.ID, err, nil)

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
		criminalCase,
	)
}

func UpdateCriminalCase(c *gin.Context) {

	var req dto.UpdateCriminalCaseRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	criminalCase, changes, err := criminalCaseService.UpdateCriminalCase(
		c.Param("id"),
		req,
	)

	LogCreate(c, constants.ActionUpdate ,&criminalCase.ID, err, changes)

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
		criminalCase,
	)
}

func GetCriminalCases(c *gin.Context) {

	search := c.Query("search")

	caseType := c.Query("case_type")

	pagination := utils.GetPagination(c)

	criminalCases, pagination, err := criminalCaseService.GetCriminalCases(
		search,
		caseType,
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
		criminalCases,
		pagination,
	)
}

func GetCriminalCaseByID(c *gin.Context) {

	id := c.Param("id")

	criminalCase, err := criminalCaseService.GetCriminalCaseByID(
		id,
	)

	LogCreate(c, constants.ActionView ,&criminalCase.ID, err, nil)

	if err != nil {

		if err.Error() == constants.NotFound("criminal case") {

			response.Error(
				c,
				http.StatusNotFound,
				err.Error(),
			)

			return
		}

		response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	response.Success(
		c,
		criminalCase,
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

	description := action + " Criminal Case"

	if recordID != nil {
		description += ": " + recordID.String()
	}

	if err != nil {

		activityLogService.LogFailure(
			userID,
			"officer_intelligence",
			action,
			description,
			err,
		)

		return
	}

	activityLogService.LogSuccess(
		userID,
		"officer_intelligence",
		action,
		recordID,
		description,
		changes,
	)
}