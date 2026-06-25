package auth

import (
	"k9-system/config"
	// "k9-system/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

import (
	userModel "k9-system/models/user"
	response "k9-system/response"
)

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type JWTClaims struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
	Email string `json:"email"`
	FullName string `json:"user_full_name"`

	jwt.RegisteredClaims
}

func Login(c *gin.Context) { // c *gin.Context contain request, response, headers, body, params

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	var user userModel.User // create empty user object

	// SELECT * FROM users WHERE email = 'admin@k9.com' LIMIT 1;
	config.DB.
		Where("email = ?", req.Email).
		First(&user)

	if user.ID.String() == "" {

		response.Error(
			c,
			http.StatusUnauthorized,
			"Invalid credentials",
		)

		return
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {

		response.Error(
			c,
			http.StatusUnauthorized,
			"Invalid credentials",
		)
		return
	}

	// Create JWT claims
	claims := JWTClaims{
		UserID: user.ID.String(),
		RoleID: user.RoleID.String(),
		Email: user.Email,
		FullName: user.FullName,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	secret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {

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

	config.DB.Save(&user) // save updated user

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token": tokenString,
	})
}