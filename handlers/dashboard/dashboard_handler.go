package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

import (
	response "k9-system/response"
	dashboardService "k9-system/services/dashboard"
)


func GetDashboard(c *gin.Context) {

	data, err := dashboardService.GetDashboard()

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
		data,
	)
}