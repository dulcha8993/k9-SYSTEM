package auth

import (
	"errors"
	"k9-system/config"
	"k9-system/constants"

	// "k9-system/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	userModel "k9-system/models/user"

	response "k9-system/response"
	activityLogService "k9-system/services/activity_log"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type JWTClaims struct {
	UserID   string `json:"user_id"`
	RoleID   string `json:"role_id"`
	Email    string `json:"email"`
	FullName string `json:"user_full_name"`

	jwt.RegisteredClaims
}

func Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	var user userModel.User

	result := config.DB.
		Where("email = ?", req.Email).
		First(&user)

	// User does not exist
	if result.Error != nil {

		err := errors.New("invalid credentials")

		LogAuthActivity(
			c,
			constants.ActionLogin,
			nil,
			err,
		)

		response.Error(
			c,
			http.StatusUnauthorized,
			"Invalid credentials",
		)

		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	); err != nil {

		LogAuthActivity(
			c,
			constants.ActionLogin,
			&user.ID,
			err,
		)

		response.Error(
			c,
			http.StatusUnauthorized,
			"Invalid credentials",
		)

		return
	}

	// Create JWT claims
	claims := JWTClaims{
		UserID:   user.ID.String(),
		RoleID:   user.RoleID.String(),
		Email:    user.Email,
		FullName: user.FullName,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}

	// Create token
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	secret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString(
		[]byte(secret),
	)

	if err != nil {

		LogAuthActivity(
			c,
			constants.ActionLogin,
			&user.ID,
			err,
		)

		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to generate token",
		)

		return
	}

	// Update last login
	now := time.Now()

	user.LastLoginAt = &now

	if err := config.DB.Save(&user).Error; err != nil {

		LogAuthActivity(
			c,
			constants.ActionLogin,
			&user.ID,
			err,
		)

		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to update last login",
		)

		return
	}

	// Successful login
	LogAuthActivity(
		c,
		constants.ActionLogin,
		&user.ID,
		nil,
	)

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"token":   tokenString,
		},
	)
}

func LogAuthActivity(
	c *gin.Context,
	action string,
	recordID *uuid.UUID,
	err error,
) {
	var userID uuid.UUID

	// During login there may not be a user_id in context yet.
	if value, exists := c.Get("user_id"); exists {
		if id, ok := value.(uuid.UUID); ok {
			userID = id
		}
	}

	description := action + " AUTH"

	if recordID != nil {
		description += ": " + recordID.String()
	}

	if err != nil {
		activityLogService.LogFailure(
			userID,
			"AUTH",
			action,
			description,
			err,
		)
		return
	}

	activityLogService.LogSuccess(
		*recordID,
		"AUTH",
		action,
		recordID,
		description,
		nil,
	)
}
