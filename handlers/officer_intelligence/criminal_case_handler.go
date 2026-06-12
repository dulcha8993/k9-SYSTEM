package officer_intelligence

import (
	// "time"

	"github.com/gin-gonic/gin"
	"net/http"
	// "gorm.io/gorm"

	// "k9-system/config"
	"k9-system/utils"
)

// import (
// 	CriminalCaseModel "k9-system/models/officer_intelligence"
// )

import (
	dto "k9-system/dto/officer_intelligence"
	response "k9-system/response"
	constants "k9-system/constants"
	criminalCaseService "k9-system/services/officer_intelligence"
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

	criminalCase, err := criminalCaseService.UpdateCriminalCase(
		c.Param("id"),
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

