package routes

import (
	"myturbogarage/helpers"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

type httpHandler func(*gin.Context) error

func check(f httpHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := f(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}

// authCheck checks for auth errors in the handler
func authCheck(f httpHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized, no token provided"})
			return
		}

		token := strings.Split(authHeader, "Bearer ")
		if len(token) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized, invalid token"})
			return
		}

		claims, err := helpers.ValidateToken(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("claims", claims)
		if err := f(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
