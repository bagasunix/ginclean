package middlewares

import (
	"strings"

	"github.com/bagasunix/ginclean/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "You are not logged in! Please log in to get access."})
			context.Abort()
			return
		}

		err := jwt.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
