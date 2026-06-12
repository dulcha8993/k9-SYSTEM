package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"k9-system/constants"
	"k9-system/response"
)

func ValidateUUIDOrAbort(
	c *gin.Context,
	id string,
	field string,
) (uuid.UUID, bool) {

	parsedUUID, err := uuid.Parse(id)

	if err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			constants.InvalidUUID(field),
		)

		return uuid.Nil, false
	}

	return parsedUUID, true
}