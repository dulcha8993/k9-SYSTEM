package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

import authHandler "k9-system/handlers/auth"

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization") //Reads Authorization Header, Validates JWT, Extracts Claims (Get user_id, role_id, email)

		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authorization header missing",
			})

			c.Abort()

			return
		}

		// Expected format:
		// Bearer TOKEN

		splitToken := strings.Split(authHeader, " ")

		if len(splitToken) != 2 {

			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid token format",
			})

			c.Abort()

			return
		}

		tokenString := splitToken[1]

		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.ParseWithClaims(
			tokenString,
			&authHandler.JWTClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			},
		)

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid token",
			})

			c.Abort()

			return
		}

		claims, ok := token.Claims.(*authHandler.JWTClaims)

		if !ok || !token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid token claims",
			})

			c.Abort()

			return
		}

		// Store user info in request context


		userUUID, err := uuid.Parse(claims.UserID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid user ID",
			})
			return
		}

		c.Set("user_id", userUUID)
		// c.Set("user_id", claims.UserID)
		c.Set("role_id", claims.RoleID)
		c.Set("email", claims.Email)
		c.Set("user_full_name", claims.FullName)

		c.Next()
	}
}