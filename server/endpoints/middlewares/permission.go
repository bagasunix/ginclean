package middlewares

import (
	"strings"

	"github.com/bagasunix/ginclean/server/domains/entities"
	"github.com/gin-gonic/gin"
)

func Permission(stArr ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := ctx.Value("authorization_payload").(entities.Account)
		low := payload.Role.Name
		isValid := false
		for _, v := range stArr {
			stArr := strings.ToLower(v)
			if strings.Contains(stArr, low) {
				isValid = true
			}
		}
		if !isValid {
			ctx.JSON(403, gin.H{"error": "You do not have permission to perform this action"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
