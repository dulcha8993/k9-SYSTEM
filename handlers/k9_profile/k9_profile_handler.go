package k9_profile

import (
	"net/http"
	// "fmt"
	"gorm.io/gorm" 

	"github.com/gin-gonic/gin"
)

import (
	utils "k9-system/utils"
)

import (
	dto "k9-system/dto/k9_profile"
	response "k9-system/response"
	k9ProfileService "k9-system/services/k9_profile"
)

func CreateK9Profile(c *gin.Context) {

	var req dto.CreateK9ProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	record, err := k9ProfileService.CreateK9Profile(
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

	response.Created(c,record,)
}

func UpdateK9Profile(c *gin.Context) {

	var req dto.UpdateK9ProfileRequest

	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	record, err := k9ProfileService.UpdateK9Profile(
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

func GetK9Profiles(c *gin.Context) {

	search := c.Query("search")

	status := c.Query("status")

	microchip := c.Query("microchip")	

	pagination := utils.GetPagination(c)

	
	k9Profiles, pagination, err := k9ProfileService.GetK9Profiles(
		search,
		status,
		microchip,
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
		k9Profiles,
		pagination,
	)
}

func GetK9ProfileByID(c *gin.Context) {

	id := c.Param("id")

	
	k9Profile, err :=  k9ProfileService.GetK9ProfileByID(
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
	response.Success(c,k9Profile,)
}