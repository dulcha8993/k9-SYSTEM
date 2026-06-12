package k9_health

import (
	"net/http"
	"gorm.io/gorm" 

	"github.com/gin-gonic/gin"
)

import (
	utils "k9-system/utils"
)

import (
	dto "k9-system/dto/k9_health"
	response "k9-system/response"
	k9HealthService "k9-system/services/k9_health"
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


	record, err := k9HealthService.UpdateHealthRecord(
		id,
		req,
	)

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