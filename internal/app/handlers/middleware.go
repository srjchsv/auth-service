package handlers

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/auth-service/internal/app/services"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			return
		}

		userId, err := services.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		log.Printf("User entered the secured zone, id:%v ", userId)
		c.Next()
	}
}
