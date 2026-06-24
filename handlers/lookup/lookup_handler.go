package lookup

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"k9-system/response"

	lookupService "k9-system/services/lookup"
)

func GetLookups(c *gin.Context) {

	data, err := lookupService.GetLookups()

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