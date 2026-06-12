package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Paginated(
	c *gin.Context,
	data interface{},
	pagination interface{},
) {

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data: data,
		Pagination: pagination,
	})
}