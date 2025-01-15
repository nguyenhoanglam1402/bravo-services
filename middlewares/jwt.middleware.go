package middlewares

import (
	"bravo-service/api/packages/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
			c.Abort()
			return
		}

		// Remove "Bearer " from the token string
		if len(tokenString) > 7 && strings.ToLower(tokenString[0:7]) == "bearer " {
			tokenString = tokenString[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}

		// Validate the token
		claims, err := helper.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}
<<<<<<< HEAD
		// Set the username from the token claims to the context
		c.Set("username", claims["username"])
		c.Set("email", claims["username"])
		c.Set("role", claims["role"])
		c.Set("uid", claims["uid"])
=======

		// Set the username from the token claims to the context
		c.Set("username", claims["username"])
		c.Set("email", claims["username"])
>>>>>>> c2516f8374928f024a808e95b22156a2ad3ad03f

		c.Next()
	}
}
