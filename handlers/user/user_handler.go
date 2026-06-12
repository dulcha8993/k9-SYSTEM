package user

import (
	// "k9-system/config"
	"net/http"

	"github.com/gin-gonic/gin" // HTTP web framework for Go (Golang)
	// "golang.org/x/crypto/bcrypt"
	// "github.com/google/uuid"
)

import (
	// accessModel "k9-system/models/access"
	// userModel "k9-system/models/user"
	utils "k9-system/utils"
)

import (
	dto "k9-system/dto/user"
	response "k9-system/response"
	// constants "k9-system/constants"
	userService "k9-system/services/user"
)

func CreateUser(c *gin.Context) {

	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	user, err := userService.CreateUser(
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
		user,
	)
}

func UpdateUser(c *gin.Context) {

	var req dto.UpdateUserRequest

	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	user, err := userService.UpdateUser(
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
		user,
	)
}

func GetUsers(c *gin.Context) {

	search := c.Query("search")

	status := c.Query("status")

	pagination := utils.GetPagination(c)

	users, pagination, err := userService.GetUsers(
		search,
		status,
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
		users,
		pagination,
	)
}


func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	userResponse, err := userService.GetUserByID(
		id,
	)

	if err != nil {

		response.Error(
			c,
			http.StatusNotFound,
			err.Error(),
		)

		return
	}

	response.Success(
		c,
		userResponse,
	)
}

